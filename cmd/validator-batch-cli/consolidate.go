package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	validatorbatchcli "github.com/pk910/validator-batch-cli"
	"github.com/pk910/validator-batch-cli/utils"
	"github.com/sirupsen/logrus"
)

type ConsolidateArgs struct {
	sourcePubkeys []string
	targetPubkey  string
}

func runConsolidate(ctx context.Context, cliArgs CliArgs, consolidateArgs ConsolidateArgs) {
	batcher, err := startBatcher(ctx, cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start batcher: %v", err)
	}

	defer batcher.Shutdown()

	if len(consolidateArgs.sourcePubkeys) == 0 {
		logrus.Fatalf("No source pubkeys provided")
	}

	if consolidateArgs.targetPubkey == "" {
		logrus.Fatalf("No target pubkey provided")
	}

	targetPubkey := common.FromHex(consolidateArgs.targetPubkey)
	if len(targetPubkey) != 48 {
		logrus.Fatalf("Invalid target pubkey: %v", targetPubkey)
	}

	sourcePubkeys := make([][]byte, len(consolidateArgs.sourcePubkeys))
	targetPubkeys := make([][]byte, len(consolidateArgs.sourcePubkeys))
	for i, pubkey := range consolidateArgs.sourcePubkeys {
		sourcePubkeys[i] = common.FromHex(pubkey)
		if len(sourcePubkeys[i]) != 48 {
			logrus.Fatalf("Invalid source pubkey: %v", pubkey)
		}

		targetPubkeys[i] = targetPubkey
	}

	// check batcher contract
	batcherDeployed, err := batcher.CheckBatcher(ctx)
	if err != nil {
		logrus.Fatalf("Failed to check batcher: %v", err)
	}

	if !batcherDeployed {
		fmt.Println("Batcher contract not deployed!")

		// check factory
		factoryDeployed, err := batcher.CheckFactory(ctx)
		if err != nil {
			logrus.Fatalf("Failed to check factory: %v", err)
		}

		if !factoryDeployed {
			logrus.Fatalf("Keyless contract factory not deployed!")

			if cliArgs.approveFactory {
				fmt.Println("Deploying factory contract...")

				tx, receipt, err := batcher.DeployFactory(ctx, cliArgs.approveFactory)
				if err != nil {
					logrus.Fatalf("Failed to deploy factory: %v", err)
				}

				fmt.Printf("Factory deployment tx hash: %v (nonce: %v)\n", tx.Hash(), tx.Nonce())
				if receipt.Status == types.ReceiptStatusFailed {
					fmt.Printf("  TX Status: reverted\n")
				} else {
					fmt.Printf("  TX Status: success\n")
				}

				fmt.Printf("  TX Fee: %v\n", utils.FormatAmount(big.NewInt(0).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))))
			} else {
				logrus.Fatalf("Keyless contract factory not deployed! Use --approve-factory to deploy, this will cost 0.01 ETH.")
			}
		}

		if cliArgs.approveBatcher {
			fmt.Println("Deploying batcher contract...")

			tx, receipt, err := batcher.DeployBatcher(ctx)
			if err != nil {
				logrus.Fatalf("Failed to deploy batcher: %v", err)
			}

			fmt.Printf("Batcher deployment tx hash: %v (nonce: %v)\n", tx.Hash(), tx.Nonce())
			if receipt.Status == types.ReceiptStatusFailed {
				fmt.Printf("  TX Status: reverted\n")
			} else {
				fmt.Printf("  TX Status: success\n")
			}

			fmt.Printf("  TX Fee: %v\n", utils.FormatAmount(big.NewInt(0).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))))

		} else {
			logrus.Fatalf("Batcher contract not deployed! Use --approve-batcher to deploy, this will cost ETH worth about 500k gas (0.0014 ETH at 3 Gweo gas price).")
		}
	}

	// check delegate
	delegate, err := batcher.GetCurrentDelegate(ctx)
	if err != nil {
		logrus.Fatalf("Failed to get current delegate: %v", err)
	}

	needDelegate := false
	if !bytes.Equal(delegate.Bytes(), validatorbatchcli.BatcherAddress.Bytes()) {
		if !cliArgs.approveDelegate {
			logrus.Fatalf("Wallet is not delegated to the batcher contract! Use --approve-delegate to delegate, this will not cost any additional ETH, but impacts your wallet behavior. You can remove the delegation via the `validator-batch-cli delegate -d` command.")
		}

		needDelegate = true
	}

	tx, err := batcher.CreateBatchConsolidateTx(ctx, sourcePubkeys, targetPubkeys, cliArgs.requestFeeLimit, needDelegate)
	if err != nil {
		logrus.Fatalf("Failed to create batch consolidate tx: %v", err)
	}

	fmt.Printf("Batch consolidate tx hash: %v (nonce: %v)\n", tx.Hash(), tx.Nonce())

	/*
		txBytes, err := tx.MarshalBinary()
		if err != nil {
			logrus.Fatalf("Failed to marshal tx: %v", err)
		}

		fmt.Printf("  TX Data: %x\n", txBytes)

		txJson, err := json.Marshal(tx)
		if err != nil {
			logrus.Fatalf("Failed to marshal tx: %v", err)
		}

		fmt.Printf("  TX JSON: %s\n", string(txJson))
	*/

	_, receipt, err := batcher.SubmitAndAwaitConfirmation(ctx, tx)
	if err != nil {
		logrus.Fatalf("Failed to submit and await confirmation: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		fmt.Printf("  TX Status: reverted\n")
	} else {
		fmt.Printf("  TX Status: success\n")
	}

	fmt.Printf("  TX Fee: %v\n", utils.FormatAmount(big.NewInt(0).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))))
}

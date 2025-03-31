package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pk910/validator-batch-cli/utils"
	"github.com/sirupsen/logrus"
)

func runDeploy(ctx context.Context, cliArgs CliArgs) {
	batcher, err := startBatcher(ctx, cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start batcher: %v", err)
	}

	defer batcher.Shutdown()

	// check factory
	factoryDeployed, err := batcher.CheckFactory(context.Background())
	if err != nil {
		logrus.Fatalf("Failed to check factory: %v", err)
	}

	if factoryDeployed {
		logrus.Info("Factory already deployed, skipping...")
	} else {
		logrus.Info("Factory not deployed, deploying...")
		tx, receipt, err := batcher.DeployFactory(context.Background(), cliArgs.approveBatcher)
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
	}

	// check batcher
	batcherDeployed, err := batcher.CheckBatcher(context.Background())
	if err != nil {
		logrus.Fatalf("Failed to check batcher: %v", err)
	}

	if batcherDeployed {
		logrus.Info("Batcher already deployed, skipping...")
	} else {
		logrus.Info("Batcher not deployed, deploying...")
		tx, receipt, err := batcher.DeployBatcher(context.Background())
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
	}

	logrus.Info("Deployment complete.")
}

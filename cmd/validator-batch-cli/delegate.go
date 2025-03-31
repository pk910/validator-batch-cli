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

type DelegateArgs struct {
	delegateAddress  string
	removeDelegation bool
}

func runDelegate(ctx context.Context, cliArgs CliArgs, delegateArgs DelegateArgs) {
	batcher, err := startBatcher(ctx, cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start batcher: %v", err)
	}

	defer batcher.Shutdown()

	currentDelegation, err := batcher.GetCurrentDelegate(ctx)
	if err != nil {
		logrus.Fatalf("Failed to get current delegate: %v", err)
	}

	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	if bytes.Equal(currentDelegation.Bytes(), zeroAddress.Bytes()) && delegateArgs.removeDelegation {
		logrus.Fatalf("Wallet is not delegated to any contract")
	}

	delegateTarget := validatorbatchcli.BatcherAddress
	if delegateArgs.removeDelegation {
		delegateTarget = zeroAddress
	} else if delegateArgs.delegateAddress != "" {
		delegateTarget = common.HexToAddress(delegateArgs.delegateAddress)
	}

	tx, receipt, err := batcher.Delegate(ctx, delegateTarget)
	if err != nil {
		logrus.Fatalf("Failed to delegate: %v", err)
	}

	fmt.Printf("Delegation tx hash: %v (nonce: %v)\n", tx.Hash(), tx.Nonce())
	if receipt.Status == types.ReceiptStatusFailed {
		fmt.Printf("  TX Status: reverted\n")
	} else {
		fmt.Printf("  TX Status: success\n")
	}

	fmt.Printf("  TX Fee: %v\n", utils.FormatAmount(big.NewInt(0).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))))

	logrus.Info("Deployment complete.")
}

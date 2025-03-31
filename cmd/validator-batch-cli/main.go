package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	validatorbatchcli "github.com/pk910/validator-batch-cli"
)

type CliArgs struct {
	verbose         bool
	config          string
	rpchosts        []string
	privkey         string
	batcherAddress  string
	baseFee         uint64
	tipFee          uint64
	approveFactory  bool
	approveBatcher  bool
	approveDelegate bool
	requestFeeLimit uint64
}

func main() {
	cliArgs := CliArgs{}
	rootCmd := &cobra.Command{
		Use:   "fundingvault",
		Short: "Funding vault CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			runMain(context.Background(), cliArgs)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&cliArgs.verbose, "verbose", "v", false, "Run the script with verbose output")
	rootCmd.PersistentFlags().StringVarP(&cliArgs.config, "config", "c", "", "The config file to use.")
	rootCmd.PersistentFlags().StringVarP(&cliArgs.privkey, "privkey", "p", "", "The private key of the wallet to send funds from.")
	rootCmd.PersistentFlags().StringSliceVarP(&cliArgs.rpchosts, "rpchost", "r", []string{}, "The RPC host to send transactions to.")
	rootCmd.PersistentFlags().StringVar(&cliArgs.batcherAddress, "batcher", "", "The address of the batcher contract.")
	rootCmd.PersistentFlags().Uint64Var(&cliArgs.baseFee, "basefee", 20, "Max fee per gas to use in claim transaction (in gwei)")
	rootCmd.PersistentFlags().Uint64Var(&cliArgs.tipFee, "tipfee", 2, "Max tip per gas to use in claim transaction (in gwei)")
	rootCmd.PersistentFlags().BoolVar(&cliArgs.approveDelegate, "approve-delegate", false, "Approve the delegate to the batcher contract.")
	rootCmd.PersistentFlags().BoolVar(&cliArgs.approveFactory, "approve-factory", false, "Approve the deployment of the contract factory.")
	rootCmd.PersistentFlags().BoolVar(&cliArgs.approveBatcher, "approve-batcher", false, "Approve the deployment of the batcher contract.")
	rootCmd.PersistentFlags().Uint64Var(&cliArgs.requestFeeLimit, "request-fee-limit", 1000000000, "Request a fee limit for the transaction (in wei).")
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy the batcher contract",
		Run: func(cmd *cobra.Command, args []string) {
			runDeploy(context.Background(), cliArgs)
		},
	}
	rootCmd.AddCommand(deployCmd)

	delegateArgs := DelegateArgs{}
	delegateCmd := &cobra.Command{
		Use:   "delegate",
		Short: "Delegate wallet to the batcher contract",
		Run: func(cmd *cobra.Command, args []string) {
			runDelegate(context.Background(), cliArgs, delegateArgs)
		},
	}
	delegateCmd.Flags().StringVarP(&delegateArgs.delegateAddress, "address", "a", "", "The address of the contract to delegate to.")
	delegateCmd.Flags().BoolVarP(&delegateArgs.removeDelegation, "remove", "d", false, "Remove active delegation from the wallet.")
	rootCmd.AddCommand(delegateCmd)

	consolidateArgs := ConsolidateArgs{}
	consolidateCmd := &cobra.Command{
		Use:   "consolidate",
		Short: "Consolidate a group of validators into a single validator",
		Run: func(cmd *cobra.Command, args []string) {
			runConsolidate(context.Background(), cliArgs, consolidateArgs)
		},
	}
	consolidateCmd.Flags().StringSliceVarP(&consolidateArgs.sourcePubkeys, "source", "s", []string{}, "The source validator pubkeys to consolidate.")
	consolidateCmd.Flags().StringVarP(&consolidateArgs.targetPubkey, "target", "t", "", "The target validator pubkey to consolidate to.")
	rootCmd.AddCommand(consolidateCmd)

	rootCmd.Execute()
}

func startBatcher(ctx context.Context, cliArgs CliArgs) (*validatorbatchcli.Batcher, error) {
	if cliArgs.verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	var config *validatorbatchcli.BatcherConfig
	if cliArgs.config != "" {
		var err error
		config, err = validatorbatchcli.LoadConfig(cliArgs.config)
		if err != nil {
			panic(err)
		}
	} else {
		config = validatorbatchcli.NewConfig()
	}

	for _, rpcHost := range strings.Split(strings.Join(cliArgs.rpchosts, ","), ",") {
		if rpcHost != "" {
			config.RpcHosts = append(config.RpcHosts, rpcHost)
		}
	}

	if cliArgs.privkey != "" {
		config.Privkey = cliArgs.privkey
	}

	if cliArgs.batcherAddress != "" {
		config.BatcherAddress = cliArgs.batcherAddress
	}

	if cliArgs.baseFee > 0 {
		config.TxBaseFee = cliArgs.baseFee
	}

	if cliArgs.tipFee > 0 {
		config.TxTipFee = cliArgs.tipFee
	}

	logger := logrus.New()
	logger.SetLevel(logrus.GetLevel())

	batcher := validatorbatchcli.NewBatcher(ctx, config, logger)
	if err := batcher.Initialize(); err != nil {
		batcher.Shutdown()
		return nil, err
	}

	return batcher, nil
}

func runMain(ctx context.Context, cliArgs CliArgs) {
	batcher, err := startBatcher(ctx, cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start batcher: %v", err)
	}

	defer batcher.Shutdown()

	fmt.Println("Validator batch actions cli-tool")
	fmt.Println("================================")

	delegateAddress, err := batcher.GetCurrentDelegate(ctx)
	if err != nil {
		logrus.Fatalf("Failed to get current delegate: %v", err)
	}

	fmt.Printf("Current delegate: %v\n", delegateAddress.String())

}

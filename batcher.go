package validatorbatchcli

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethpandaops/spamoor/spamoor"
	"github.com/ethpandaops/spamoor/txbuilder"
	"github.com/ethpandaops/spamoor/utils"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"

	"github.com/pk910/validator-batch-cli/contract"
)

var (
	FactoryAddress   = common.HexToAddress("0x7a0d94f55792c434d74a40883c6ed8545e406d12")
	FactoryDeployer  = common.HexToAddress("0x4c8D290a1B368ac4728d83a9e8321fC3af2b39b1")
	FactoryDeployTx  = common.FromHex("0xf87e8085174876e800830186a08080ad601f80600e600039806000f350fe60003681823780368234f58015156014578182fd5b80825250506014600cf31ba022222222222222222222222222222222222222222222222222222222222222222a022222222222222222222222222222222222222222222222222222222222222222")
	BatcherAddress   = crypto.CreateAddress2(FactoryAddress, [32]byte{}, crypto.Keccak256(common.FromHex(string(contract.ContractMetaData.Bin))))
	BatcherDeployGas = uint64(650000)
)

type Batcher struct {
	ctx        context.Context
	cancel     context.CancelFunc
	Config     *BatcherConfig
	Logger     *logrus.Logger
	clientPool *spamoor.ClientPool
	txpool     *txbuilder.TxPool
	rootWallet *txbuilder.Wallet
}

func NewBatcher(ctx context.Context, config *BatcherConfig, logger *logrus.Logger) *Batcher {
	ctx, cancel := context.WithCancel(ctx)
	return &Batcher{
		ctx:    ctx,
		cancel: cancel,
		Config: config,
		Logger: logger,
	}
}

func (b *Batcher) Initialize() error {
	// init client pool
	b.clientPool = spamoor.NewClientPool(b.ctx, b.Config.RpcHosts, b.Logger)
	err := b.clientPool.PrepareClients()
	if err != nil {
		return err
	}

	// prepare txpool
	b.txpool = txbuilder.NewTxPool(&txbuilder.TxPoolOptions{
		GetClientFn: func(index int, random bool) *txbuilder.Client {
			mode := spamoor.SelectClientByIndex
			if random {
				mode = spamoor.SelectClientRandom
			}
			return b.clientPool.GetClient(mode, index)
		},
		GetClientCountFn: func() int {
			return len(b.clientPool.GetAllGoodClients())
		},
	})

	// load root wallet
	if b.Config.Privkey == "" {
		return fmt.Errorf("privkey is not set")
	}
	rootWallet, err := txbuilder.NewWallet(b.Config.Privkey)
	if err != nil {
		return err
	}
	b.rootWallet = rootWallet

	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	err = client.UpdateWallet(b.ctx, b.rootWallet)
	if err != nil {
		return err
	}

	b.Logger.Infof(
		"initialized root wallet (addr: %v balance: %v ETH, nonce: %v)",
		rootWallet.GetAddress().String(),
		utils.WeiToEther(uint256.MustFromBig(rootWallet.GetBalance())).Uint64(),
		rootWallet.GetNonce(),
	)

	return nil
}

func (b *Batcher) Shutdown() {
	b.cancel()
}

func (b *Batcher) SubmitAndAwaitConfirmation(ctx context.Context, tx *types.Transaction) (*types.Transaction, *types.Receipt, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return nil, nil, fmt.Errorf("no client available")
	}

	var receipt *types.Receipt
	var txErr error

	wg := sync.WaitGroup{}
	wg.Add(1)
	err := b.txpool.SendTransaction(ctx, b.rootWallet, tx, &txbuilder.SendTransactionOptions{
		Client:              client,
		MaxRebroadcasts:     10,
		RebroadcastInterval: 30 * time.Second,
		OnConfirm: func(tx *types.Transaction, rec *types.Receipt, err2 error) {
			defer func() {
				wg.Done()
			}()

			if err2 != nil {
				txErr = err2
			}

			receipt = rec
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send tx: %v", err)
	}

	wg.Wait()
	if txErr != nil {
		return tx, nil, fmt.Errorf("failed to wait for tx receipt: %v", txErr)
	}

	return tx, receipt, nil
}

func (b *Batcher) CreateDelegationTx(txMetadata *txbuilder.TxMetadata, delegateAddress common.Address) (*types.Transaction, error) {
	txNonce := b.rootWallet.GetNextNonce()
	authNonce := b.rootWallet.GetNextNonce()

	authorization := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(b.rootWallet.GetChainId().Uint64()),
		Address: delegateAddress,
		Nonce:   authNonce,
	}

	signedAuthorization, err := types.SignSetCode(b.rootWallet.GetPrivateKey(), authorization)
	if err != nil {
		return nil, fmt.Errorf("could not sign set code authorization: %v", err)
	}

	txMetadata.AuthList = []types.SetCodeAuthorization{signedAuthorization}

	txData, err := txbuilder.SetCodeTx(txMetadata)
	if err != nil {
		return nil, fmt.Errorf("failed to create set code tx: %v", err)
	}

	txData.ChainID = uint256.MustFromBig(b.rootWallet.GetChainId())
	txData.Nonce = txNonce
	tx := types.NewTx(txData)
	tx, err = types.SignTx(tx, types.LatestSignerForChainID(b.rootWallet.GetChainId()), b.rootWallet.GetPrivateKey())
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx: %v", err)
	}

	return tx, nil
}

func (b *Batcher) GetCurrentDelegate(ctx context.Context) (common.Address, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return common.Address{}, fmt.Errorf("no client available")
	}

	code, err := client.GetEthClient().CodeAt(ctx, b.rootWallet.GetAddress(), nil)
	if err != nil {
		return common.Address{}, err
	}

	if len(code) == 0 {
		return common.Address{}, nil
	}

	// check if code is a eip 7702 delegate
	if !bytes.HasPrefix(code, []byte("\xef\x01\x00")) {
		return common.Address{}, nil
	}

	// get delegate address
	delegateAddress := common.BytesToAddress(code[3:])

	return delegateAddress, nil
}

func (b *Batcher) CheckFactory(ctx context.Context) (bool, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return false, fmt.Errorf("no client available")
	}

	code, err := client.GetEthClient().CodeAt(ctx, FactoryAddress, nil)
	if err != nil {
		return false, err
	}

	if len(code) == 0 {
		return false, nil
	}

	return true, nil
}

func (b *Batcher) DeployFactory(ctx context.Context, approveTopUp bool) (*types.Transaction, *types.Receipt, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return nil, nil, fmt.Errorf("no client available")
	}

	// init deployer wallet
	deployerWallet, err := txbuilder.NewWallet("")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create deployer wallet: %v", err)
	}

	deployerWallet.SetAddress(FactoryDeployer)

	err = client.UpdateWallet(ctx, deployerWallet)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to update deployer wallet: %v", err)
	}

	if deployerWallet.GetNonce() > 0 {
		return nil, nil, fmt.Errorf("deployer wallet already has a nonce")
	}

	// check if deployer has at least 0.01 ETH
	if deployerWallet.GetBalance().Cmp(big.NewInt(10000000000000000)) < 0 {
		if !approveTopUp {
			return nil, nil, fmt.Errorf("deployer has less than 0.01 ETH")
		}

		// top up
		txData, err := txbuilder.DynFeeTx(&txbuilder.TxMetadata{
			GasFeeCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxBaseFee) * 1e9)),
			GasTipCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxTipFee) * 1e9)),
			Gas:       25000,
			To:        &FactoryDeployer,
			Value:     uint256.MustFromBig(big.NewInt(10000000000000000)),
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create topup tx: %v", err)
		}

		tx, err := b.rootWallet.BuildDynamicFeeTx(txData)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to sign topup tx: %v", err)
		}

		_, _, err = b.SubmitAndAwaitConfirmation(ctx, tx)
		if err != nil {
			return nil, nil, fmt.Errorf("topup tx failed: %v", err)
		}

		// refresh balance
		err = client.UpdateWallet(ctx, deployerWallet)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to update deployer wallet: %v", err)
		}

		if deployerWallet.GetBalance().Cmp(big.NewInt(10000000000000000)) < 0 {
			return nil, nil, fmt.Errorf("failed to top up deployer")
		}
	}

	factoryTx := &types.Transaction{}
	err = factoryTx.UnmarshalBinary(FactoryDeployTx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal factory deploy tx: %v", err)
	}

	_, receipt, err := b.SubmitAndAwaitConfirmation(ctx, factoryTx)
	if err != nil {
		return nil, nil, fmt.Errorf("factory deploy tx failed: %v", err)
	}

	return factoryTx, receipt, nil
}

func (b *Batcher) CheckBatcher(ctx context.Context) (bool, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return false, fmt.Errorf("no client available")
	}

	logrus.Infof("Checking batcher at address: %v", BatcherAddress.String())

	code, err := client.GetEthClient().CodeAt(ctx, BatcherAddress, nil)
	if err != nil {
		return false, err
	}

	if len(code) == 0 {
		return false, nil
	}

	return true, nil
}

func (b *Batcher) DeployBatcher(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	batcherCode := common.FromHex(string(contract.ContractMetaData.Bin))
	txData, err := txbuilder.DynFeeTx(&txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxBaseFee) * 1e9)),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxTipFee) * 1e9)),
		Gas:       BatcherDeployGas,
		To:        &FactoryAddress,
		Value:     uint256.NewInt(0),
		Data:      batcherCode,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create deploy tx: %v", err)
	}

	tx, err := b.rootWallet.BuildDynamicFeeTx(txData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign tx: %v", err)
	}

	return b.SubmitAndAwaitConfirmation(ctx, tx)
}

func (b *Batcher) Delegate(ctx context.Context, delegateAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	walletAddress := b.rootWallet.GetAddress()
	txMetadata := &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxBaseFee) * 1e9)),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxTipFee) * 1e9)),
		Gas:       50000,
		To:        &walletAddress,
		Value:     uint256.NewInt(0),
	}

	tx, err := b.CreateDelegationTx(txMetadata, delegateAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create delegation tx: %v", err)
	}

	return b.SubmitAndAwaitConfirmation(ctx, tx)
}

func (b *Batcher) CreateBatchConsolidateTx(ctx context.Context, sourcePubkeys [][]byte, targetPubkeys [][]byte, feeLimit uint64, withDelegate bool) (*types.Transaction, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return nil, fmt.Errorf("no client available")
	}

	walletAddress := b.rootWallet.GetAddress()
	txMetadata := &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxBaseFee) * 1e9)),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxTipFee) * 1e9)),
		Gas:       1000000,
		To:        &walletAddress,
		Value:     uint256.NewInt(0),
	}

	batchContract, err := contract.NewContract(walletAddress, client.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create batch contract: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(b.rootWallet.GetPrivateKey(), b.rootWallet.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	transactor.Context = ctx
	transactor.From = walletAddress
	transactor.GasLimit = txMetadata.Gas
	transactor.GasTipCap = txMetadata.GasTipCap.ToBig()
	transactor.GasFeeCap = txMetadata.GasFeeCap.ToBig()
	transactor.Value = txMetadata.Value.ToBig()
	transactor.NoSend = true

	data := make([][]byte, len(sourcePubkeys))
	for i, sourcePubkey := range sourcePubkeys {
		data[i] = make([]byte, 96)
		copy(data[i][:48], sourcePubkey)
		copy(data[i][48:], targetPubkeys[i])
	}

	boundTx, err := batchContract.BatchConsolidate(transactor, data, big.NewInt(int64(feeLimit)))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch consolidate tx: %v", err)
	}

	txMetadata.Data = boundTx.Data()

	var signedTx *types.Transaction
	if withDelegate {
		signedTx, err = b.CreateDelegationTx(txMetadata, BatcherAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to create delegation tx: %v", err)
		}
	} else {
		txData, err := txbuilder.DynFeeTx(txMetadata)
		if err != nil {
			return nil, fmt.Errorf("failed to create tx: %v", err)
		}

		signedTx, err = b.rootWallet.BuildDynamicFeeTx(txData)
		if err != nil {
			return nil, fmt.Errorf("failed to create tx: %v", err)
		}
	}

	return signedTx, nil
}

func (b *Batcher) CreateBatchWithdrawTx(ctx context.Context, sourcePubkeys [][]byte, amounts []uint64, feeLimit uint64, withDelegate bool) (*types.Transaction, error) {
	client := b.clientPool.GetClient(spamoor.SelectClientRandom, 0)
	if client == nil {
		return nil, fmt.Errorf("no client available")
	}

	walletAddress := b.rootWallet.GetAddress()
	txMetadata := &txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxBaseFee) * 1e9)),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(b.Config.TxTipFee) * 1e9)),
		Gas:       1000000,
		To:        &walletAddress,
		Value:     uint256.NewInt(0),
	}

	batchContract, err := contract.NewContract(walletAddress, client.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create batch contract: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(b.rootWallet.GetPrivateKey(), b.rootWallet.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	transactor.Context = ctx
	transactor.From = walletAddress
	transactor.GasLimit = txMetadata.Gas
	transactor.GasTipCap = txMetadata.GasTipCap.ToBig()
	transactor.GasFeeCap = txMetadata.GasFeeCap.ToBig()
	transactor.Value = txMetadata.Value.ToBig()
	transactor.NoSend = true

	data := make([][]byte, len(sourcePubkeys))
	for i, sourcePubkey := range sourcePubkeys {
		data[i] = make([]byte, 56)
		copy(data[i][:48], sourcePubkey)

		amount := big.NewInt(0).SetUint64(amounts[i])
		amountBytes := amount.FillBytes(make([]byte, 8))
		copy(data[i][48:], amountBytes)
	}

	boundTx, err := batchContract.BatchConsolidate(transactor, data, big.NewInt(int64(feeLimit)))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch consolidate tx: %v", err)
	}

	txMetadata.Data = boundTx.Data()

	var signedTx *types.Transaction
	if withDelegate {
		signedTx, err = b.CreateDelegationTx(txMetadata, walletAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to create delegation tx: %v", err)
		}
	} else {
		txData, err := txbuilder.DynFeeTx(txMetadata)
		if err != nil {
			return nil, fmt.Errorf("failed to create tx: %v", err)
		}

		signedTx, err = b.rootWallet.BuildDynamicFeeTx(txData)
		if err != nil {
			return nil, fmt.Errorf("failed to create tx: %v", err)
		}
	}

	return signedTx, nil
}

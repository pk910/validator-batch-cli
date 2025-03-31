package validatorbatchcli

import (
	"os"

	"gopkg.in/yaml.v3"
)

type BatcherConfig struct {
	Privkey  string   `yaml:"privkey"`
	RpcHosts []string `yaml:"rpchosts"`

	BatcherAddress string `yaml:"batcherAddress"`
	TxBaseFee      uint64 `yaml:"txBaseFee"`
	TxTipFee       uint64 `yaml:"txTipFee"`

	ApproveCreateFactory bool `yaml:"approveCreateFactory"`
	ApproveCreateBatcher bool `yaml:"approveCreateBatcher"`
	Interactive          bool `yaml:"interactive"`
}

func NewConfig() *BatcherConfig {
	return &BatcherConfig{
		TxBaseFee: 20000000000, // 20 gwei
		TxTipFee:  1000000000,  // 1 gwei
	}
}

func LoadConfig(configFile string) (*BatcherConfig, error) {
	config := NewConfig()
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

package utils

import (
	"fmt"
	"math/big"
)

func FormatAmount(amount *big.Int) string {
	value := new(big.Int).Set(amount)
	gwei, _ := value.Div(value, big.NewInt(1e9)).Float64()
	if gwei > 0 && gwei < 100000 {
		return fmt.Sprintf("%.0f gwei", gwei)
	}

	decimalVal := fmt.Sprintf("%.4f", gwei/1e9)
	for i := len(decimalVal) - 1; i >= 0; i-- {
		if decimalVal[i] == '0' {
			decimalVal = decimalVal[:i]
		} else {
			break
		}
	}

	if decimalVal[len(decimalVal)-1] == '.' {
		decimalVal = decimalVal[:len(decimalVal)-1]
	}

	return fmt.Sprintf("%s ETH", decimalVal)
}

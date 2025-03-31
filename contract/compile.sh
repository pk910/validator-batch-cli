#!/bin/sh

solc --standard-json ./ValidatorActionBatcher.input.json | jq '.contracts["contracts/ValidatorActionBatcher.sol"].ValidatorActionBatcher.abi' > ValidatorActionBatcher.abi
solc --standard-json ./ValidatorActionBatcher.input.json | jq -r '.contracts["contracts/ValidatorActionBatcher.sol"].ValidatorActionBatcher.evm.bytecode.object' > ValidatorActionBatcher.bin
abigen --bin=./ValidatorActionBatcher.bin --abi=./ValidatorActionBatcher.abi --pkg=contract --out=ValidatorActionBatcher.go

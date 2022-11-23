

clean:
	rm -rf ./pkg/staking
	npx hardhat clean

compile:
	npx hardhat compile

build-go:
	mkdir -p ./pkg/staking
	solc --abi contracts/Staking.sol -o ./pkg/staking/ --overwrite
	solc --bin contracts/Staking.sol -o ./pkg/staking/ --overwrite
	abigen --bin=./build/Staking.bin --abi=./pkg/staking/Staking.abi --pkg=staking --out=./pkg/staking/Staking.go

build: clean compile build-go


.PHONY: compile clean build-go
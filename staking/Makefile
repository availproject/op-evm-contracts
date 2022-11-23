

clean:
	rm -rf ./build
	npx hardhat clean

compile:
	npx hardhat compile

build-go:
	mkdir ./build
	solc --abi contracts/Staking.sol -o build/ --overwrite
	solc --bin contracts/Staking.sol -o build/ --overwrite
	abigen --bin=./build/Staking.bin --abi=./build/Staking.abi --pkg=staking --out=./build/Staking.go

build: clean compile build-go


.PHONY: compile clean build-go
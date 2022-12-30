.PHONY: compile clean build-go

build: clean compile build-go

clean:
	rm -rf ./pkg/testtoken
	npx hardhat clean

compile:
	npx hardhat compile

build-go:
	mkdir -p ./pkg/testtoken
	solc --abi --bin --include-path ./node_modules --base-path . --overwrite -o ./pkg/testtoken/ contracts/TestToken.sol
	abigen --bin=./pkg/testtoken/TestToken.bin --abi=./pkg/testtoken/TestToken.abi --pkg=testtoken --out=./pkg/testtoken/TestToken.go


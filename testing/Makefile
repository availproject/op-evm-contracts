.PHONY: compile clean build-go

build: clean compile build-go

clean:
	rm -rf ./pkg/testtoken
	npx hardhat clean

compile:
	npx hardhat compile

build-go-testtoken:
	mkdir -p ./pkg/testtoken
	solc --abi --bin --include-path ./node_modules --base-path . --overwrite -o ./pkg/testtoken/ contracts/TestToken.sol
	abigen --bin=./pkg/testtoken/TestToken.bin --abi=./pkg/testtoken/TestToken.abi --pkg=testtoken --out=./pkg/testtoken/TestToken.go

build-go-testnft:
	mkdir -p ./pkg/testnft
	solc --abi --bin --include-path ./node_modules --base-path . --overwrite -o ./pkg/testnft/ contracts/TestNFT.sol
	abigen --bin=./pkg/testnft/TestNFT.bin --abi=./pkg/testnft/TestNFT.abi --pkg=testnft --out=./pkg/testnft/TestNFT.go

build-go: build-go-testtoken build-go-testnft

# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy-nft.ts
```

## Setup

```shell
cp .env.example .env
```

To set `AVAIL_SL_URL` you need to use the load balancer address from the terraform aws deployment.

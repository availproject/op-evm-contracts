# Optimistic EVM Rollup Contracts
This repository contains the Solidity smart contracts for the [Optimistic EVM Rollup](https://github.com/availproject/op-evm) solution.

The repository is structured with the `staking` and `testing` directories, 
each containing their respective contracts and configurations.

## Staking
The `staking` directory holds the staking contracts, 
which provide APIs for watchtowers and sequencers to stake their tokens. 
By staking their tokens, these participants can actively contribute to the consensus of the chain. 
Additionally, the staking contract offers APIs for slashing malicious nodes and initiating dispute 
resolutions against malicious actors.

## Testing
The `testing` directory contains various testing contracts, including a sample `ERC20` token and a `NFT` token. 
These contracts primarily serve the purpose of end-to-end tests and demonstrations. In the testing environment, 
you will also find scripts for deploying the assets and minting tokens.

## Go Bindings
To enhance the integration of these contracts with Go projects using Go modules, 
the repository includes the generated Go source code for each contract. 
The `abigen` tool was utilized to generate this code, simplifying the process of 
importing the contracts in any Go project.

Feel free to explore the repository and leverage these contracts to build and test Optimistic EVM Rollup solutions. 
Should you encounter any issues or have questions, please refer to the project documentation or open an issue on GitHub.

Thank you for your interest in the Optimistic EVM Rollup.
Should you have any questions or need further assistance, please feel free to reach out. Happy coding!

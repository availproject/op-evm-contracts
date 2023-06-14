import "dotenv/config";
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

/** @type import('hardhat/config').HardhatUserConfig */
const config: HardhatUserConfig = {
  solidity: "0.8.17",
  defaultNetwork: "availsl",
  networks: {
    availsl: {
      url: process.env.AVAIL_SL_URL,
      chainId: 100,
      gas: 21000,
      accounts: [`${process.env.ACC_PRIVATE_KEY}`]
    }
  },
};

export default config;
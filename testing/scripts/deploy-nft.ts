import { ethers } from "hardhat";
import {TestNFT} from "../typechain-types";

async function main() {
  let testNFT: TestNFT
  if (process.env.TEST_NFT_CONTRACT_ADDR) {
    testNFT = await ethers.getContractAt("TestNFT", process.env.TEST_NFT_CONTRACT_ADDR)
  } else {
    const testNFTFactory = await ethers.getContractFactory("TestNFT");
    testNFT = await testNFTFactory.deploy();
    await testNFT.waitForDeployment();
  }

  const addr = await testNFT.getAddress()

  await testNFT.mintNFT()
  const tokenUri = await testNFT.tokenURI(0)
  const tokenCounter = await testNFT.getTokenCounter()

  console.log(`TestNFT deployed to ${addr}; token uri is: ${tokenUri}; tokenCounter is: ${tokenCounter}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

import { ethers } from "hardhat";

async function main() {
  const TestNFT = await ethers.getContractFactory("TestNFT");

  const testNFT = await TestNFT.deploy();

  await testNFT.waitForDeployment();

  const addr = await testNFT.getAddress()

  const tokenUri = await testNFT.tokenURI(0)

  console.log(`TestNFT deployed to ${addr}; token uri is: ${tokenUri}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

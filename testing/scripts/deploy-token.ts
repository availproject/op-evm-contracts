import { ethers } from "hardhat";
import {TestToken} from "../typechain-types";

async function main() {
  let testToken: TestToken
  if (process.env.TEST_TOKEN_CONTRACT_ADDR) {
    testToken = await ethers.getContractAt("TestToken", process.env.TEST_TOKEN_CONTRACT_ADDR)
  } else {
    const testTokenFactory = await ethers.getContractFactory("TestToken");
    testToken = await testTokenFactory.deploy();
    await testToken.waitForDeployment();
  }

  const [owner] = await ethers.getSigners()

  await testToken.mint(owner.address, '1000000000000000000')
  const addr = await testToken.getAddress()

  console.log(`TestToken deployed to ${addr}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

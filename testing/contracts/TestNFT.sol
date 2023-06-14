
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract TestNFT is ERC721 {
    string public constant TOKEN_URI = "ipfs://bafybeig37ioir76s7mg5oobetncojcm3c3hxasyd4rvid4jqhy4gkaheg4/?filename=0-PUG.json";
    uint private s_tokenCounter = 0;
    constructor() ERC721("Test Avail NFT", "TANFT") {
    }

    function mintNFT() public returns(uint) {
        _safeMint(msg.sender, s_tokenCounter);
        s_tokenCounter++;
        return s_tokenCounter;
    }

    function getTokenCounter() public returns(uint) {
        return s_tokenCounter;
    }

    function tokenURI(uint tokenId) public view override returns (string memory) {
        return TOKEN_URI;
    }
}

pragma solidity ^0.5.0;

import "OpenZeppelin/token/ERC721Full.sol";
import "OpenZeppelin/token//ERC721Holder.sol";

contract PicBlock is ERC721Full, ERC721Holder {
//可铸造的ERC721代币
  constructor (
  )
  ERC721Full("Pic Block Token", "PBT") public {
  }
  function mint(uint256 _tokenId, string memory _tokenUri) public {
    _mint(msg.sender, _tokenId);
    _setTokenURI(_tokenId, _tokenUri);
  }
}


/// @title 墨客联盟链BAAS ERC721合约
/// @author 野驴阿保机
/// @notice 这个合约是对标准ERC721合约的扩展

pragma solidity ^0.5.0;

import "OpenZeppelin/token/ERC721Full.sol";
import "OpenZeppelin/token//ERC721Holder.sol";

contract BaasBlock is ERC721Full, ERC721Holder {
  // 这是每一个721token的固定属性，铸造的时候设定
  mapping(uint256 => string) public property;

  function() external payable {}

  constructor () ERC721Full("Baas 721 Token", "B7T") public {
  }

  event tokenMinted(uint256 indexed tokenId, string property, string tokenUri);
  event memoAdded(uint256 indexed tokenId, string memo);
  event transactionMemo(uint256 indexed tokenId, string memo);

  /// @notice 设置一个token的属性
  /// @param _tokenId 唯一的通证id
  /// @param _property 属性是一串字符串，从一个结构转化而来。
  function setProperty(uint256 _tokenId, string memory _property) public {
    address owner = ownerOf(_tokenId);
    require(msg.sender == owner, "only contract owner");

    property[_tokenId] = _property;
  }
  
  /// @notice 读取一个token的属性
  /// @param _tokenId 唯一的通证id
  function getProperty(uint256 _tokenId) public view returns (string memory) {
    return property[_tokenId];
  }

  /// @notice 为给定的token添加一条memo
  /// @param _tokenId 唯一的通证id
  /// @param _memo 记录一些附加的信息
  function addMemo(uint256 _tokenId, string memory _memo) public {
    // 不做检查，如果是空的字符串，就空的字符串
    // require(bytes(_memo).length > 0, "should not be empty memo")
    emit memoAdded(_tokenId, _memo);
  }

  function transferWithMemo(address _from, address _to, uint256 _tokenId, string memory _memo) public {
    transferFrom(_from, _to, _tokenId);
    if (bytes(_memo).length > 0) {
      emit transactionMemo(_tokenId, _memo);
    }
  }

  function mint(uint256 _tokenId, string memory _property, string memory _tokenUri) public {
    _mint(msg.sender, _tokenId);

    // 这里不对property做判断 - mint的时候，如果不需要property, 则使用""，
    // 这样这个token总是有个property. 因此getProperty不会出错。
    property[_tokenId] = _property;

    // 这里设置token的metadata, 这里也不使用baseUrl
    _setTokenURI(_tokenId, _tokenUri);

    emit tokenMinted(_tokenId, _property, _tokenUri);
  }
}


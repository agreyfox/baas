/// @title 墨客联盟链BAAS ERC721合约
/// @author 野驴阿保机
/// @notice 这个合约是对标准ERC721合约的扩展

pragma solidity ^0.5.0;
import "OpenZeppelin/token/ERC721/ERC721Full.sol";
import "OpenZeppelin/token/ERC721/ERC721Holder.sol";

//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC721/ERC721Full.sol";
//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC721/ERC721Holder.sol";

contract ERC721Base is ERC721Full, ERC721Holder {
  // 这是每一个721token的固定属性，铸造的时候设定
  mapping(uint256 => string) public property;

  function() external payable {}

  constructor (
    string memory name, //代币名称
    string memory symbol //代币缩写
  ) ERC721Full(name, symbol) public {
  }

  event tokenMinted(uint256 indexed tokenId, string tokenProperty, string tokenUri);
  event memoAdded(uint256 indexed tokenId, string memo);
  event transactionMemo(uint256 indexed tokenId, string memo);

  /// @notice 设置一个token的属性
  /// @param tokenId 唯一的通证id
  /// @param tokenProperty 属性是一串字符串，从一个结构转化而来。
  function setProperty(uint256 tokenId, string memory tokenProperty) public {
    address owner = ownerOf(tokenId);
    require(msg.sender == owner, "only contract owner");

    property[tokenId] = tokenProperty;
  }
  
  /// @notice 读取一个token的属性
  /// @param tokenId 唯一的通证id
  function getProperty(uint256 tokenId) public view returns (string memory) {
    return property[tokenId];
  }

  /// @notice 为给定的token添加一条memo
  /// @param tokenId 唯一的通证id
  /// @param memo 记录一些附加的信息
  function addMemo(uint256 tokenId, string memory memo) public {
    // 不做检查，如果是空的字符串，就空的字符串
    // require(bytes(_memo).length > 0, "should not be empty memo")
    emit memoAdded(tokenId, memo);
  }

  function transferWithMemo(address from, address to, uint256 tokenId, string memory memo) public {
    transferFrom(from, to, tokenId);
    if (bytes(memo).length > 0) {
      emit transactionMemo(tokenId, memo);
    }
  }

  function mint(uint256 tokenId, string memory tokenProperty, string memory tokenUri) public {
    _mint(msg.sender, tokenId);

    // 这里不对tokenProperty做判断 - mint的时候，如果不需要tokenProperty, 则使用""，
    // 这样这个token总是有个tokenProperty. 因此getProperty不会出错。
    property[tokenId] = tokenProperty;

    // 这里设置token的metadata, 这里也不使用baseUrl
    _setTokenURI(tokenId, tokenUri);

    emit tokenMinted(tokenId, tokenProperty, tokenUri);
  }
}


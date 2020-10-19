pragma solidity 0.5.0;

import "OpenZeppelin/token/ERC20/ERC20.sol";
import "OpenZeppelin/token/ERC20/ERC20Detailed.sol";
import "OpenZeppelin/token/ERC20/ERC20Burnable.sol";
import "OpenZeppelin/token/ERC20/ERC20Pausable.sol";
import "OpenZeppelin/token/ERC20/ERC20Mintable.sol";

//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC20/ERC20.sol";
//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC20/ERC20Detailed.sol";
//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC20/ERC20Burnable.sol";
//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC20/ERC20Pausable.sol";
//import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v2.5.0/contracts/token/ERC20/ERC20Mintable.sol";

//多功能ERC20代币,可销毁,可暂停
contract ERC20WithBurnPauseMint is
    ERC20,
    ERC20Detailed,
    ERC20Burnable,
    ERC20Pausable,
    ERC20Mintable
{
    constructor(
        string memory name, //代币名称
        string memory symbol, //代币缩写
        uint8 decimals, //精度
        uint256 totalSupply //发行总量
    ) public ERC20Detailed(name, symbol, decimals) {
        _mint(msg.sender, totalSupply * (10**uint256(decimals)));
    }

    // 添加交易memo功能, 这里调用了标准的transfer功能，transfer已经考虑了whenNotPaused的情况
    event transactionMemo(string memo);

    function transferWithMemo(address recipient, uint256 amount, string memory memo) public {
        if (transfer(recipient, amount)) {
            if (bytes(memo).length > 0) {
                emit transactionMemo(memo);
            }
        }
    }

    function transferFromWithMemo(address sender, address recipient, uint256 amount, string memory memo) public {
        if (transferFrom(sender, recipient, amount)) {
            if (bytes(memo).length > 0) {
                emit transactionMemo(memo);
            }
        }
    }

}

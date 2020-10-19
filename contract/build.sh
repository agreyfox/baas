rm -rf 721abi
../../solc-static-linux --abi -o 721abi baas721Base.sol 
abigen --pkg baas721 --abi 721abi/ERC721Base.abi --out baas721Base.go

rm -rf abi
../../solc-static-linux --abi -o abi baas721.sol 
abigen --pkg baas721 --abi abi/BaasBlock.abi --out baas721.go

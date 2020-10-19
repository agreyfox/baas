rm -rf 20abi

./solc --abi -o 20abi baas20Standard.sol 
abigen --pkg baas20 --abi 20abi/ERC20Standard.abi --out baas20Base.go

rm -rf 20abi
./solc  --abi -o 20abi baas20BurnPause.sol 
abigen --pkg baas20 --abi 20abi/ERC20WithBurnPause.abi --out baas20BurnPause.go


rm -rf 20abi
./solc --abi -o 20abi baas20BurnPauseMint.sol 
abigen --pkg baas20 --abi 20abi/ERC20WithBurnPauseMint.abi --out baas20BurnPauseMint.go


rm -rf 20abi
./solc --abi -o 20abi  baas20BurnPauseMintCap.sol 
abigen --pkg baas20 --abi 20abi/ERC20MultiFunction.abi --out baas20BurnPauseMintCap.go


package baas

import (
	"context"
	"time"
)

var (
	PageSize = 25
)

const (
	BlockEngine            = "http://127.0.0.1:8545"
	GasMax                 = 90000000
	GasLimit               = 21000
	GasMsgLimit            = 1800000
	GasERC20Limit          = 500000
	GasERC721Limit         = 1000000
	GasERC721TransferLimit = 100000
)

var (
	BlockChainServer        = BlockEngine
	DelplyContractGas int64 = 5000000
	OriginTokenOP     int64 = 545228
	ERC20TokenOP      int64 = 100000
	ERC20TokenSend    int64 = 500000
	ERC721TokenOP     int64 = 500000
	ERC721TokenMint   int64 = 700000
	ERC721TokenSend   int64 = 700000
)

// ApplicationService defines the business logic for dealing with all aspects of an application.
type BlockService interface {
	Create(ctx context.Context, n *NewBAASUser) (*BAASUser, error)
	Import(ctx context.Context, n *NewBAASUser) (*BAASUser, error)
	ChangePassword(ctx context.Context, userId, oldPassword, newPassword string) error
	DeleteBAASUser(ctx context.Context, id string) error
	UpdateBAASUser(ctx context.Context, u *UpdateBAASUser) (*BAASUser, error)
	GetKey(ctx context.Context, id, password, ciper string) (string, error)
	DeleteKey(ctx context.Context, id, passwd, applicationid string) error
	GetAddress(ctx context.Context, id, password string) (string, error)
	RecoverKey(ctx context.Context, id string) (string, error)
	GetBalance(ctx context.Context, addr string) (string, error)
	SendToken(ctx context.Context, userid, password, toAddr, value, msg, encode string, gas int64) (string, error)
	WriteMsg(ctx context.Context, addr, password, toAddr, msg string, gas int64) (string, error)
	ReadMsg(ctx context.Context, hash, encode string) (string, string, error)
	GetTxByHash(ctx context.Context, hash string) (string, error)
	//===================erc721 ====================================
	GetErc721Info(ctx context.Context, addr string) (map[string]interface{}, error)
	GetErc721Balance(ctx context.Context, addr, contractAddr string) (string, error)
	GetErc721TotalSupply(ctx context.Context, addr string) (string, error)
	CreateErc721Token(ctx context.Context, userid, password, addr, tokenid, meta, property string, gas int64) (string, error)
	SetErc721TokenProperty(ctx context.Context, userid, password, addr, tokenid, property string, gas int64) (string, error)
	GetErc721TokenProperty(ctx context.Context, addr, tokenid string) (string, error)
	SendErc721Token(ctx context.Context, addr, password, contractaddr, tokeid, memo, targetuser string, gas int64) (string, error)
	GetSendErc721TokenMemo(ctx context.Context, contractaddr, tokeid, hash string) (string, error)
	GetErc721MetaData(ctx context.Context, addr, tokenid string) (string, error)
	AddErc721TokenMemo(ctx context.Context, user, password, contract, tokenid, memo string, gas int64) (string, error)
	GetErc721TokenMemoList(ctx context.Context, contractaddr, tokenid string) (string, error)
	GetUserErc721TokenList(ctx context.Context, userid, contractaddr string) (string, error)
	CreateERC721Contract(ctx context.Context, userid, password, name, symbol, class string) (string, string, error) //user create new ERC721 contranct with type
	//===================erc20 ====================================
	DeployERC20Contract(ctx context.Context, userid, password, name, symbol, class string, totalsupply uint64, decimal uint8, capacity uint64) (string, string, error) //user create new ERC20 contranct with type
	GetErc20TotalSupply(ctx context.Context, addr string) (string, error)
	GetErc20Decimal(ctx context.Context, addr string) (string, error)
	GetErc20Info(ctx context.Context, addr string) (map[string]interface{}, error)
	GetErc20Balance(ctx context.Context, userid, addr string) (string, error)
	SendErc20Token(ctx context.Context, addr, password, toAddr, contract, memo string, value float64, gas int64) (string, error)
	ApproveErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error)
	AllowanceErc20(ctx context.Context, addr, toAddr, contract string) (string, error)
	IncreaseAllowanceErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error)
	DecresaseAllowanceErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error)
	GetErc20TxMemo(ctx context.Context, hash string) (string, error)
	BurnErc20(ctx context.Context, addr, password, contract string, value float64, gas int64) (string, error)
	PauseErc20(ctx context.Context, addr, password, contract string, value bool, gas int64) (string, error)
	MintErc20(ctx context.Context, addr, password, contract string, value float64, gas int64) (string, error)
	GetErc20PauseStatus(ctx context.Context, contract string) (string, error)
	TransferFromErc20(ctx context.Context, user, password, fromaddr, toAddr, contract, memo string, value float64, gas int64) (string, error)
	// ==================++Erc20 analysis
	GetERC20TxByUserAddress(ctx context.Context, userid, contract, page, pagesize string) (string, error)
	GetERC20TxList(ctx context.Context, contract, page, pagesize string) (string, error)
	//CreateErc20Token(ctx context.Context, addr string) (string, error)
	// The following the some statistics functions
	//GetErc721Property(ctx context.Context, addr, tokenid string) (string, error)
	//===================analysis and utility ====================================
	GetApplicationUsers(ctx context.Context, id string, page, limit int) ([]*BAASUser, error)
	GetTx(ctx context.Context, addr string) (*BlockTx, error)
	GetBlockNumber(ctx context.Context) (string, error)
	GetAddressFromPK(ctx context.Context, pk string) (string, error)
	GetKeyStoreFromPK(ctx context.Context, pk, password string) (string, error)
	GetPKFromKeyStore(ctx context.Context, keystore, password string) (string, error)
	//FileBlobStorage(engine, accessKey, secretKey, region, endpoint string) (FileBlobStorage, error)
	DecryptKey(origin, cipherFromWeb, cipherFromDb, rv, salt string) (string, error)
	GetTxByUserAddress(ctx context.Context, usrid, page, size string) (string, error)                     // query backend system for address tx record
	GetPeerToPeerTxByUserAddress(ctx context.Context, usrid, targetid, page, size string) (string, error) // query backend system for address tx record
	GetErc721TxList(ctx context.Context, addr, page, size string) (string, error)
	GetErc721TokenTxList(ctx context.Context, contract, tokenid, page, size string) (string, error)
	GetErc721TxListByUser(ctx context.Context, userid, contract, page, size string) (string, error)
	GetErc721TokenTxListByUser(ctx context.Context, userid, contract, tokenid, page, size string) (string, error)
}

// blockStorage handles communication with the database for handling block.
type BlockStorage interface {
	Store(ctx context.Context, n *NewBAASUser) (*BAASUser, error)
	UpdatePassword(ctx context.Context, id, old, new string) error
	GetKey(ctx context.Context, id, password string) (string, string, string, string, error)
	DeleteKey(ctx context.Context, id, password string) error
	GetAddress(ctx context.Context, id, password string) (string, error)
	// address, origpk, rv,cipher,salt
	GetAddressByService(ctx context.Context, id, password string) (string, string, string, string, string, error)
	//Delete(ctx context.Context, id string) error
	GetApplicationId(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, u *UpdateBAASUser) (*BAASUser, error)
	ApplicationUsers(ctx context.Context, id string, page, limit int) ([]*BAASUser, error)
}

type Block struct {
	ID               string
	Name             string
	Description      string
	StorageAccessKey string
	StorageSecretKey string
	StorageBucket    string
	StorageEndpoint  string
	StorageRegion    string
	StorageEngine    string
	DeliveryURL      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UpdateBAASUser struct {
	ID          string
	Email       string
	UserName    string
	Description string
	Cmd         string
	Address     string
	PrivateKey  string
	Password    string
}

type NewBAASUser struct {
	Email            string
	Name             string
	Description      string
	StorageAccessKey string
	StorageSecretKey string
	StorageBucket    string
	ApplicationID    string
	Address          string
	PrivateKey       string
	CipherText       string
	Rv               string
	Salt             string
	Password         string
	Key              string
}

type BAASUser struct {
	ID               string
	Email            string
	Name             string
	Description      string
	StorageAccessKey string
	StorageSecretKey string
	StorageBucket    string
	ApplicationID    string
	Address          string
	Password         string
	PrivateKey       string
	Salt             string
	CipherText       string
	Rv               string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BlockTx struct {
	Type   string
	Filter string
	Data   []interface{}
}

package baas

import (
	"context"
	"time"
)

var ()

const (
	BlockEngine            = "http://127.0.0.1:8545"
	GasLimit               = 21000
	GasMsgLimit            = 1800000
	GasERC20Limit          = 100000
	GasERC721Limit         = 1000000
	GasERC721TransferLimit = 100000
)

var (
	BlockChainServer = BlockEngine
)

// ApplicationService defines the business logic for dealing with all aspects of an application.
type BlockService interface {
	Create(ctx context.Context, n *NewBAASUser) (*BAASUser, error)
	ChangePassword(ctx context.Context, userId, oldPassword, newPassword string) error
	DeleteBAASUser(ctx context.Context, id string) error
	UpdateBAASUser(ctx context.Context, u *UpdateBAASUser) (*BAASUser, error)
	GetKey(ctx context.Context, id, password, ciper string) (string, error)
	DeleteKey(ctx context.Context, id, passwd string) error
	GetAddress(ctx context.Context, id, password string) (string, error)
	RecoverKey(ctx context.Context, id string) (string, error)
	GetBalance(ctx context.Context, addr string) (string, error)
	SendToken(ctx context.Context, userid, password, toAddr, value string) (string, error)
	WriteMsg(ctx context.Context, addr, password, toAddr, msg string) (string, error)
	ReadMsg(ctx context.Context, hash string) (string, string, error)
	GetTxByHash(ctx context.Context, hash string) (string, error)
	GetErc20Balance(ctx context.Context, addr string) (string, error)
	SendErc20Token(ctx context.Context, addr string) (string, error)
	CreateErc20Token(ctx context.Context, addr string) (string, error)
	GetErc721Balance(ctx context.Context, addr string) (string, error)
	GetErc721TotalSupply(ctx context.Context, addr string) (string, error)
	CreateErc721Token(ctx context.Context, addr string) (string, error)
	SendErc721Token(ctx context.Context, addr string) (string, error)
	GetErc721MetaData(ctx context.Context, addr string) (string, error)
	GetTx(ctx context.Context, addr string) (*BlockTx, error)
	//FileBlobStorage(engine, accessKey, secretKey, region, endpoint string) (FileBlobStorage, error)
	DecryptKey(origin, cipherFromWeb, cipherFromDb, rv, salt string) (string, error)
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
	Update(ctx context.Context, u *UpdateBAASUser) (*BAASUser, error)
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

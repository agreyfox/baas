package block

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agreyfox/baas"

	"github.com/agreyfox/baas/cmd/baasd"
	"github.com/agreyfox/baas/storm"
	"github.com/rs/zerolog"
)

var (
	StormClient *storm.EthRPC
	BigHelper   = big.NewInt(10)
	Wei         = BigHelper.Exp(BigHelper, big.NewInt(18), nil)
)

type Service struct {
	BlockStorage       baas.BlockStorage
	ApplicationService baas.ApplicationService
	EncryptionService  baas.EncryptionService
	GmService          baas.EncryptionService
	UsageService       baas.UsageService
	Log                zerolog.Logger
}

// block service first connect storm block chain service
func InitBlockService() {
	conf := baasd.GetSystemConfig()
	if len(conf.Blockchain.Connection) == 0 {
		fmt.Errorf("BlockChain backend service is not config!,system exit.")
		os.Exit(-1)
	}
	fmt.Println("Trying to connect Blocachain service...", conf.Blockchain.Connection)
	StormClient = storm.New(conf.Blockchain.Connection, storm.WithDebug(true))

	version, err := StormClient.Web3ClientVersion()
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	fmt.Println("Blockchain backend service connected! \n Version ", version)
}

func ExtractCipher(c string) ([]byte, []byte) {
	buf := []byte(c)
	if len(c) == 0 {
		fmt.Println("Warning, No input cipher , will use system default gm key and iv")
		config := baasd.GetSystemConfig()
		return []byte(config.Security.GMIV), []byte(config.Security.GMKey)
	}
	if len(c) < 16 {
		newbuf := baas.PKCS5Padding(buf, 15)
		return buf[0:1], newbuf
	} else {
		return buf[0:1], buf[len(c)-15 : len(c)]
	}
}

func (s *Service) Create(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	fmt.Printf("Call to create the new baas user %s with application id :%s\n", n.Email, n.ApplicationID)
	add, pk, err := StormClient.NewAddress()
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate new address!")
		return nil, err
	}
	n.Address = add

	salt, _ := baas.GenerateSalt()
	ciper := n.CipherText
	if len(n.CipherText) == 0 {
		config := baasd.GetSystemConfig()
		ciper = config.Security.GMKey
	}
	rv, partCiper := ExtractCipher(ciper)
	realCiper := append(rv, partCiper...)
	s.GmService.Config(realCiper, salt)
	fmt.Println("now to encrypt:", time.Now())
	//fmt.Println("Input text:", pk)

	skey, err := s.GmService.EncryptToString([]byte(pk))
	if err != nil {
		s.Log.Error().Err(err).Msg("EncryptToString error")
		return nil, err
	}
	fmt.Println("Done:", time.Now())
	//fmt.Printf("Input ciper text:%s\n", skey)

	/* mm, err := s.GmService.DecryptString(skey)
	fmt.Println("descypt result:", string(mm)) */

	n.PrivateKey = skey
	n.Salt = base64.StdEncoding.EncodeToString(salt)
	n.CipherText = string(partCiper)
	n.Rv = string(rv)

	return s.BlockStorage.Store(ctx, n)
}

//change password with email oldpass, and newpass
func (s *Service) ChangePassword(ctx context.Context, email, oldpass, newpass string) error {
	return s.BlockStorage.UpdatePassword(ctx, email, oldpass, newpass)
}

func (s *Service) DeleteBAASUser(ctx context.Context, id string) error {
	return nil
}

func (s *Service) UpdateBAASUser(ctx context.Context, u *baas.UpdateBAASUser) (*baas.BAASUser, error) {
	return &baas.BAASUser{
		Name:             "",
		Description:      "",
		StorageAccessKey: "",
		StorageSecretKey: "",
		StorageBucket:    "",
		ApplicationID:    "",
		PrivateKey:       "",
		Password:         "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}, nil
}

func (s *Service) GetKey(ctx context.Context, id, password, ciperText string) (string, error) {
	ke, ciper, rv, salt, err := s.BlockStorage.GetKey(ctx, id, password)
	if err == nil {
		var realCiper []byte
		ciperbuf := []byte(ciper)
		if len(ciperText) > 0 {
			mbuf := []byte(ciperText)
			realrv := mbuf[:1]
			realCiper = append(realrv, ciperbuf...)
		} else {
			realCiper = append([]byte(rv), ciperbuf...)
		}
		realsalt, _ := base64.StdEncoding.DecodeString(salt)
		s.GmService.Config(realCiper, realsalt)
		realkey, err := s.GmService.DecryptString(ke)
		if err != nil {
			return "", errors.New("Internal Error")
		}

		fmt.Println(string(realkey), err)
		return string(realkey), nil
	} else {
		return "", err
	}

}
func (s *Service) DeleteKey(ctx context.Context, id, password string) error {
	return s.BlockStorage.DeleteKey(ctx, id, password)
}
func (s *Service) GetAddress(ctx context.Context, id, password string) (string, error) {
	ke, err := s.BlockStorage.GetAddress(ctx, id, password)
	if err == nil {
		return ke, nil
	} else {
		return "", errors.New("Address Not found")
	}

}
func (s *Service) RecoverKey(ctx context.Context, id string) (string, error) { return "", nil }
func (s *Service) GetBalance(ctx context.Context, addr string) (string, error) {
	var err error
	address := addr
	if !strings.HasPrefix(addr, "0x") {
		address, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, addr, "")
		if err != nil {
			s.Log.Err(err).Msg("User id is not found")
			return "", err
		}
	}
	nm, err := StormClient.EthGetBalance(address, "latest")
	v := new(big.Int).Div(&nm, StormClient.Eth1())
	return v.String(), nil

}

func (s *Service) DecryptKey(origin, cipherFromWeb, cipherFromDb, rv, salt string) (string, error) {

	var realCiper []byte
	//ciperbuf := []byte(cipherFromWeb)
	if len(cipherFromWeb) > 0 {
		mbuf := []byte(cipherFromWeb)
		realrv := mbuf[:1]
		realCiper = append(realrv, cipherFromDb...)
	} else {
		realCiper = append([]byte(rv), cipherFromDb...)
	}
	realsalt, _ := base64.StdEncoding.DecodeString(salt)
	s.GmService.Config(realCiper, realsalt)
	realkey, err := s.GmService.DecryptString(origin)
	if err != nil {
		return "", errors.New("Internal Encryption Error")
	}

	fmt.Println(string(realkey), err)
	return string(realkey), nil

}

func (s *Service) SendToken(ctx context.Context, addr, password, toAddr, value string) (string, error) {
	var err error
	from := addr
	to := toAddr
	var frompk, rv, salt, cipher string
	if !strings.HasPrefix(addr, "0x") {
		from, frompk, rv, cipher, salt, err = s.BlockStorage.GetAddressByService(ctx, addr, password)
		if err != nil {
			s.Log.Err(err).Msg("User id is not found")
			return "", err
		}
		frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
		if err != nil {
			return "", err
		}
	}
	if !strings.HasPrefix(toAddr, "0x") {
		to, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, toAddr, "")
		if err != nil {
			s.Log.Err(err).Msg("Target User id is not found")
			return "", err
		}
	}
	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		s.Log.Err(err).Msg("Value  format is not correct")
		return "", err
	}

	amount := FloatToBigInt(vv)
	//fmt.Println(storm.BigToHex(vv))
	tx := storm.T{
		From:     from,
		To:       to,
		Gas:      baas.GasLimit + 3000,
		GasPrice: big.NewInt(5000000000),
		Value:    amount, //big.NewInt(1000000000000000000),
		//Data:     "0x1111",
		Nonce: nounce,
	}
	//fmt.Println(tx)
	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
	}
	// blocknm, err := StormClient.EthSendTransaction(tx)

	return t.TransactionHash, err
}

func (s *Service) WriteMsg(ctx context.Context, addr, password, toAddr, msg string) (string, error) {
	var err error
	from := addr
	to := toAddr
	var frompk, rv, salt, cipher string
	if !strings.HasPrefix(addr, "0x") {
		from, frompk, rv, cipher, salt, err = s.BlockStorage.GetAddressByService(ctx, addr, password)
		if err != nil {
			s.Log.Err(err).Msg("User id is not found")
			return "", err
		}
		frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
		if err != nil {
			return "", err
		}
	}
	if !strings.HasPrefix(toAddr, "0x") {
		to, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, toAddr, "")
		if err != nil {
			s.Log.Err(err).Msg("Target User id is not found")
			return "", err
		}
	}
	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if f < 0.0001 {
		return "", baas.ErrBaasNotEnoughMoney
	}
	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)

	amount := FloatToBigInt(vv)

	tx := storm.T{
		From:     from,
		To:       to,
		Gas:      baas.GasMsgLimit,
		GasPrice: big.NewInt(5000000000),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     msg,
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
	}
	// blocknm, err := StormClient.EthSendTransaction(tx)

	return t.TransactionHash, err

}
func (s *Service) ReadMsg(ctx context.Context, hash string) (string, string, error) {
	var tm string
	tx, err := StormClient.EthGetTransactionByHash(hash)
	if err != nil {
		fmt.Println("error is get hash:", hash)
		return "", "", err
	}
	//fmt.Println(tx)
	retstr, _ := hex.DecodeString(tx.Input[2:]) //remove 0x
	if *tx.BlockNumber > 0 {
		block, err := StormClient.EthGetBlockByNumber(*tx.BlockNumber, false)
		if err == nil {
			ttt := time.Unix(int64(block.Timestamp), 0)
			tm = fmt.Sprint(ttt)
		}
	}
	return string(retstr), tm, nil
}

// encaspulate tx to json string and return
func (s *Service) GetTxByHash(ctx context.Context, hash string) (string, error) {
	tx, err := StormClient.EthGetTransactionByHash(hash)
	if err != nil {
		return "", err
	}
	dd, err := tx.MarshalJSON()
	return string(dd), err
}
func (s *Service) GetErc20Balance(ctx context.Context, addr string) (string, error) { return "", nil }

func (s *Service) SendErc20Token(ctx context.Context, addr string) (string, error)   { return "", nil }
func (s *Service) CreateErc20Token(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetErc721Balance(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetErc721TotalSupply(ctx context.Context, addr string) (string, error) {
	return "", nil
}
func (s *Service) CreateErc721Token(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) SendErc721Token(ctx context.Context, addr string) (string, error)   { return "", nil }
func (s *Service) GetErc721MetaData(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetTx(ctx context.Context, addr string) (*baas.BlockTx, error) {
	return &baas.BlockTx{
		Type:   "",
		Filter: "",
	}, nil
}

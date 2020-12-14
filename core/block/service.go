package block

import (
	"context"
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agreyfox/baas"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/sha3"

	"github.com/agreyfox/baas/cmd/baasd"
	"github.com/agreyfox/baas/storm"
	"github.com/rs/zerolog"
)

var (
	StormClient     *storm.EthRPC
	BigHelper       = big.NewInt(10)
	Wei             = BigHelper.Exp(BigHelper, big.NewInt(18), nil)
	Abi             abi.ABI
	GasCheckMinimum = 0.001
	GasPrice        int64
)

var (
	BackEndGetAddressTx      = "/api/v1/wallets/%s/tx"
	BackEndPeerToPeerTx      = "/api/v1/wallets/%s/%s/txfilter"
	BackEndPeerToERC20Tx     = "/api/v1/wallets/%s/%s/tx"
	BackEndPeerToERC721Tx    = "/api/v1/wallets/%s/%s/tx"
	BackEndPeerERC721TokenTx = "/api/v1/erc721/%s/%s/%s/tokentx"
)

type Service struct {
	BlockStorage       baas.BlockStorage
	ApplicationService baas.ApplicationService
	EncryptionService  baas.EncryptionService
	GmService          baas.EncryptionService
	UsageService       baas.UsageService
	Log                zerolog.Logger
}

func init() {
	var err error
	Abi, err = abi.JSON(strings.NewReader(getBaas721ABI()))
	if err != nil {
		fmt.Println("721 abi parse failed,please check ")
		return
	}
	//fmt.Println("System load 721 ABI success!")
}

// block service first connect storm block chain service
func InitBlockService() {
	conf := baasd.GetSystemConfig()
	if len(conf.Blockchain.Connection) == 0 {
		fmt.Errorf("BlockChain backend service is not config!,system exit.")
		os.Exit(-1)
	}
	fmt.Println("Trying to connect Blocachain service...", conf.Blockchain.Connection)
	StormClient = storm.New(conf.Blockchain.Connection, conf.Blockchain.NID, storm.WithDebug(true))

	version, err := StormClient.Web3ClientVersion()
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	GasPrice = int64(conf.Blockchain.GasPrice) // 设置gasPrice from configure file

	fmt.Println("Blockchain backend service connected! \n Version ", version)
	fmt.Printf("Blockchain backend service gas price set to %d!\n", GasPrice)
	BackEndGetAddressTx = conf.Blockchain.APIServer + BackEndGetAddressTx           // "/api/v1/wallets/%s/tx"
	BackEndPeerToPeerTx = conf.Blockchain.APIServer + BackEndPeerToPeerTx           // "/api/v1/wallets/%s/%s/txfilter"
	BackEndPeerToERC20Tx = conf.Blockchain.APIServer + BackEndPeerToERC20Tx         // "/api/v1/wallets/%s/%s/tx"
	BackEndPeerToERC721Tx = conf.Blockchain.APIServer + BackEndPeerToERC721Tx       // "/api/v1/wallets/%s/%s/tx"
	BackEndPeerERC721TokenTx = conf.Blockchain.APIServer + BackEndPeerERC721TokenTx //"/api/v1/erc721/%s/%s/%s/tokentx"

	NetworkID = conf.Blockchain.NID //设置网络ID

}

//return two array byte, first is 8 byte, second is 12 byte
func ExtractCipher(c string) ([]byte, []byte) {
	buf := []byte(c)
	if len(buf) == 0 {
		fmt.Println("Warning, No input cipher , will use system default gm key and iv")
		config := baasd.GetSystemConfig()
		return []byte(config.Security.GMIV), []byte(config.Security.GMKey)
	}
	if len(buf) < 8 {
		fmt.Println("Warning, cipher length is to small ,will use system default gm key and iv")
		config := baasd.GetSystemConfig()
		return []byte(config.Security.GMIV), []byte(config.Security.GMKey)
	}
	if len(buf) < 12 {
		//newbuf := baas.PKCS5Padding(buf, 12)
		return buf[0:8], buf
	} else {
		return buf[0:8], buf
	}
}

//return prev 16 byte
func md516(str []byte) []byte {
	//data := []byte(str)
	var target []byte
	if len(str) > 4 {
		target = str[0:4]
	} else {
		target = str
	}
	has := md5.Sum(target)
	//md5str := fmt.Sprintf("%x", has)
	return has[:]
}

func (s *Service) Create(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	fmt.Printf("Call to create the new baas user %s with application id :%s\n", n.Email, n.ApplicationID)
	config := baasd.GetSystemConfig()
	add, pk, err := StormClient.NewAddress()
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate new address!")
		return nil, err
	}
	n.Address = add

	salt, _ := baas.GenerateSalt()
	ciper := n.CipherText
	if len(n.CipherText) == 0 {
		//config := baasd.GetSystemConfig()
		ciper = config.Security.GMKey
	}
	rvorig, partCiper := ExtractCipher(ciper)
	//rv := rvorig) //change to 4 byte

	rvorig = append(rvorig, partCiper...)

	s.GmService.Config(md516(partCiper), salt)
	defer s.GmService.Config([]byte(config.Security.GMKey), []byte(config.Security.GMIV))
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
	n.Rv = string(rvorig)

	updatedUsages := &baas.UpdateUsage{
		ApplicationID: n.ApplicationID,
		UserCount:     1,
	}

	if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
		// should I fail the request
		s.Log.Error().Err(err).Msg("failed to update Create user  usage")
	}

	return s.BlockStorage.Store(ctx, n)
}

func (s *Service) Import(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	fmt.Printf("Call to Import private key to  the new baas user %s with application id :%s\n", n.Email, n.ApplicationID)
	config := baasd.GetSystemConfig()
	ppk := n.Key
	if ppk[0] == '0' && ppk[1] == 'x' {
		ppk = ppk[2:]
	}
	add := StormClient.NewAddressFromPK(ppk)

	n.Address = add

	salt, _ := baas.GenerateSalt()
	ciper := n.CipherText
	if len(n.CipherText) == 0 {
		//config := baasd.GetSystemConfig()
		ciper = config.Security.GMKey
	}
	rvorig, partCiper := ExtractCipher(ciper)
	//rv := rvorig) //change to 4 byte

	rvorig = append(rvorig, partCiper...)

	s.GmService.Config(md516(partCiper), salt)
	defer s.GmService.Config([]byte(config.Security.GMKey), []byte(config.Security.GMIV))
	fmt.Println("now to encrypt:", time.Now())

	skey, err := s.GmService.EncryptToString([]byte(ppk))
	if err != nil {
		s.Log.Error().Err(err).Msg("EncryptToString error")
		return nil, err
	}
	fmt.Println("Done:", time.Now())

	n.PrivateKey = skey
	n.Salt = base64.StdEncoding.EncodeToString(salt)
	n.CipherText = string(partCiper)
	n.Rv = string(rvorig)

	updatedUsages := &baas.UpdateUsage{
		ApplicationID: n.ApplicationID,
		UserCount:     1,
	}

	if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
		// should I fail the request
		s.Log.Error().Err(err).Msg("failed to update Create user  usage")
	}

	return s.BlockStorage.Store(ctx, n)
}

func (s *Service) GetAddressFromPK(ctx context.Context, pk string) (string, error) {
	privatekey := pk
	if pk[0] == '0' && pk[1] == 'x' {
		privatekey = privatekey[2:]
	}
	fmt.Printf("Call to transfer  private key to  address:%s\n", privatekey)

	add := StormClient.NewAddressFromPK(privatekey)

	return add, nil

}

func (s *Service) GetKeyStoreFromPK(ctx context.Context, pk, password string) (string, error) {
	privatekey := pk
	if pk[0] == '0' && pk[1] == 'x' {
		privatekey = privatekey[2:]
	}
	fmt.Printf("Call to transfer  private key to  address:%s\n", privatekey)
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privatekey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	addr, err := s.GetAddressFromPK(ctx, privatekey)
	if err != nil {
		return "", err
	}
	var ac accounts.Account
	if ks.HasAddress(common.HexToAddress(addr)) {
		a := accounts.Account{
			Address: common.HexToAddress(addr),
		}
		ac, err = ks.Find(a)
		if err != nil {
			return "", err
		}

	} else {
		ac, err := ks.ImportECDSA(&e, password)

		if err != nil {
			fmt.Println(err)
			return "", err
		}
		fmt.Print(ac.URL)
	}
	keyjson, err := ioutil.ReadFile(ac.URL.Path)

	return string(keyjson), nil

}

func (s *Service) GetPKFromKeyStore(ctx context.Context, ks, password string) (string, error) {

	fmt.Printf("Call to transfer keystore to private key:%s\n", ks)

	key, err := keystore.DecryptKey([]byte(ks), password)
	if err != nil {
		fmt.Println("json key failed to decrypt: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	pkstr := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println(pkstr)

	return pkstr, nil

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
	ke, _, _, salt, err := s.BlockStorage.GetKey(ctx, id, password)
	if err == nil {
		var realCiper []byte
		//ciperbuf := []byte(ciper)
		if len(ciperText) > 0 {
			mbuf := []byte(ciperText)
			//rrv, _ := ExtractCipher(ciperText)
			realCiper = mbuf // append(rrv, ciperbuf...)
		} else {
			//realCiper = append([]byte(rv), ciperbuf...)
			return "", baas.ErrBaasCipherTextRequired
		}
		realsalt, _ := base64.StdEncoding.DecodeString(salt)
		s.GmService.Config(md516(realCiper), realsalt)
		realkey, err := s.GmService.DecryptString(ke)
		if err != nil {
			return "", errors.New("Internal Error")
		}

		//fmt.Println(string(realkey), err)
		return string(realkey), nil
	} else {
		return "", err
	}

}

//remove a key
func (s *Service) DeleteKey(ctx context.Context, id, password, applicationid string) error {
	err := s.BlockStorage.DeleteKey(ctx, id, password)

	if err == nil {
		updatedUsages := &baas.UpdateUsage{
			ApplicationID: applicationid,
			UserCount:     -1,
		}

		if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
			// should I fail the request
			s.Log.Error().Err(err).Msg("failed to update Create user  usage")
		}
	}
	return err

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

	fz := new(big.Float).SetInt(&nm)
	fm := new(big.Float).SetInt(StormClient.Eth1())
	//v := new(big.Int).Div(&nm, StormClient.Eth1())
	v := new(big.Float).Quo(fz, fm)

	return v.String(), nil

}

func (s *Service) DecryptKey(origin, cipherFromWeb, cipherFromDb, rv, salt string) (string, error) {

	var realCiper []byte
	//ciperbuf := []byte(cipherFromWeb)
	if len(cipherFromWeb) > 0 {
		//mbuf := []byte(cipherFromWeb)
		_, entirekey := ExtractCipher(cipherFromWeb)
		realCiper = entirekey //append(rrv, cipherFromDb...)
	} else {
		realCiper = []byte(cipherFromDb) //append([]byte(rv), cipherFromDb...)
	}
	realsalt, _ := base64.StdEncoding.DecodeString(salt)
	s.GmService.Config(md516(realCiper), realsalt)
	realkey, err := s.GmService.DecryptString(origin)
	if err != nil {
		return "", errors.New("Internal Encryption Error")
	}

	fmt.Println(string(realkey), err)
	return string(realkey), nil

}

func (s *Service) SendToken(ctx context.Context, addr, password, toAddr, value, msg, encode string, gas int64) (string, error) {
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
	var datappayload string
	var datalen int
	if encode == "utf8" {
		if len(msg) > 0 {
			datappayload = msg
			datalen = len(datappayload)
		} else {
			datappayload = ""
			datalen = 0
		}
	} else if encode == "hex" {
		if len(msg) > 0 {

			//dt := common.HexToHash(msg)
			//dt, _ := hex.DecodeString(msg)
			ma := common.Hex2Bytes(msg)
			//dt := common.HexToAddress(msg)
			//mm := UnPaddingStr(dt.Bytes())
			//fmt.Println(mm, ma)
			datappayload = string(ma) //string(dt) //common.Bytes2Hex(dt)
			//datappayload = string(dt.Bytes())
			datalen = len(datappayload)
		}
	}
	s.Log.Log().Msgf("Total data length is %d\n", datalen)
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.OriginTokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(datalen) * 12
	tx := storm.T{
		From:     from,
		To:       to,
		Gas:      int(maxGas), //baas.GasLimit + 3000 + datalen*6,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     datappayload,
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx   usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}

		return t.TransactionHash, ee
	} else {
		return "", ee
	}
	// blocknm, err := StormClient.EthSendTransaction(tx)

}

func (s *Service) WriteMsg(ctx context.Context, addr, password, toAddr, msg string, gas int64) (string, error) {
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
	if GasPrice > 0 && f < GasCheckMinimum {
		return "", baas.ErrBaasNotEnoughMoney
	}
	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)

	amount := FloatToBigInt(vv)
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.OriginTokenOP
	}
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(msg)) * 12
	tx := storm.T{
		From:     from,
		To:       to,
		Gas:      int(maxGas),
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     msg,
		Nonce:    nounce,
	}
	//fmt.Println(tx)

	t, err := SignAndSendTx(tx, frompk)

	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx   usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}

		return t.TransactionHash, ee
	} else {
		return "", ee
	}

}

/// read msg from blockchain
func (s *Service) ReadMsg(ctx context.Context, hash, encode string) (string, string, error) {
	var tm string
	tx, err := StormClient.EthGetTransactionByHash(hash)
	if err != nil {
		fmt.Println("error is get hash:", hash)
		return "", "", err
	}
	//fmt.Println(tx)
	var retstr []byte
	if encode == "utf8" {
		retstr, _ = hex.DecodeString(tx.Input[2:]) //remove 0x
	} else if encode == "hex" {
		retstr = []byte(tx.Input)
	}

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

// encaspulate tx to json string and return
func (s *Service) GetTxByUserAddress(ctx context.Context, usrid, page, size string) (string, error) {
	var err error
	address := usrid
	if !strings.HasPrefix(address, "0x") {
		address, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, usrid, "")
		if err != nil {
			s.Log.Err(err).Msg("User id is not found")
			return "", err
		}
	}
	url := fmt.Sprintf(BackEndGetAddressTx, address)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend url ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

// encaspulate tx to json string and return
func (s *Service) GetPeerToPeerTxByUserAddress(ctx context.Context, usrid, targetid, page, size string) (string, error) {
	var err error
	address := usrid
	toAddress := targetid
	if !strings.HasPrefix(address, "0x") {
		address, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, usrid, "")
		if err != nil {
			s.Log.Err(err).Msg("From User id is not found")
			return "", err
		}
	}
	if !strings.HasPrefix(toAddress, "0x") {
		toAddress, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, targetid, "")
		if err != nil {
			s.Log.Err(err).Msg("Target User id is not found")
			return "", err
		}
	}
	url := fmt.Sprintf(BackEndPeerToPeerTx, address, toAddress)
	url += fmt.Sprintf("?page=%s", page)
	//fmt.Println("backend url ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

// return erc20 a user tx list
func (s *Service) GetERC20TxByUserAddress(ctx context.Context, userid, contract, page, pagesize string) (string, error) {
	var err error
	address := userid

	if !strings.HasPrefix(address, "0x") {
		address, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, userid, "")
		if err != nil {
			s.Log.Err(err).Msg("From User id is not found")
			return "", err
		}
	}

	url := fmt.Sprintf(BackEndPeerToERC20Tx, address, contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend url for user call erc20 tx: ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

func (s *Service) GetERC20TxList(ctx context.Context, contract, page, pagesize string) (string, error) {
	var err error
	//address := contract

	url := fmt.Sprintf(BackEndPeerToERC20Tx, "*", contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend url for all erc20  txs: ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}
func (s *Service) GetErc721Balance(ctx context.Context, addr, contractAddr string) (string, error) {
	address, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, addr, "")
	if err != nil {
		s.Log.Err(err).Msg("From User id is not found")
		return "", err
	}

	m, exists := Abi.Methods["balanceOf"]
	if exists {
		fmt.Println("Function is expected", m)
	}

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(address, &m.Inputs[0])
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: address,
		To:   contractAddr,
		Data: hexutil.Encode(dataEncodefor721Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}

	//var balance []*big.Int
	var returns = common.FromHex(tx)
	out, err := Abi.Unpack("balanceOf", returns) // modify by api change 2020/10/27
	//fmt.Println(balance)

	return fmt.Sprintf("%v", out[0]), err

}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func (s *Service) GetErc721TotalSupply(ctx context.Context, addr string) (string, error) {

	//datamethod, err := sha3.New([]byte("totalSupply()"))
	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	//hash.Write([]byte{0xcc})
	hash.Write([]byte("totalSupply()"))
	buf = hash.Sum(nil)

	//fmt.Println(buf)

	callpara := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}

	//var balance *big.Int
	var returns = common.FromHex(tx)
	out, err := Abi.Unpack("totalSupply", returns) /// modified at 2020/10/27

	return fmt.Sprintf("%v", out[0]), err

}

func (s *Service) CreateErc721Token(ctx context.Context, userid, password, contract, tokenid, meta, property string, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, userid, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	fmt.Println(from, " try to mint new token at ", contract)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)
	tid, err := strconv.ParseUint(tokenid, 10, 0)
	if err != nil {
		return "", err
	}
	if tid == 0 {
		return "", errors.New("tokenid is not correct!")
	}
	m, exists := Abi.Methods["mint"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0]) //tokenid
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(property, &m.Inputs[1]) //meta

	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(meta, &m.Inputs[2]) //property

	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 2, m.Inputs[2].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	//fmt.Println(dataEncodefor721Call)
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC721TokenMint
	}
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(meta)+len(property)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas),
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, userid)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", userid)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx mint  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}

		return t.TransactionHash, ee
	} else {
		return "", ee
	}
	return "", nil
}

func (s *Service) SetErc721TokenProperty(ctx context.Context, userid, password, contract, tokenid, property string, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, userid, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	fmt.Println(from, " try to set  new token property at ", contract, tokenid)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)
	tid, err := strconv.ParseUint(tokenid, 10, 0)
	if err != nil {
		return "", err
	}
	if tid == 0 {
		return "", errors.New("tokenid is not correct!")
	}
	m, exists := Abi.Methods["setProperty"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0]) //tokenid
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(property, &m.Inputs[1]) //property

	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC721TokenOP
	}
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(property)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC721Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("Transaction timeout!")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, userid)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", userid)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx mint  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}

		return t.TransactionHash, ee
	} else {
		return "", ee
	}
	//return "", nil
}

func (s *Service) GetErc721Info(ctx context.Context, addr string) (map[string]interface{}, error) {

	supply, err := s.GetErc721TotalSupply(ctx, addr)

	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	hash.Write([]byte("name()"))
	buf = hash.Sum(nil)

	callpara := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return map[string]interface{}{}, err
	}

	nameb := new(string)
	var returns = common.FromHex(tx)
	out, err := Abi.Unpack("name", returns)
	*nameb = fmt.Sprint(out[0]) // modified 2020/10/27

	hashsymbol := sha3.NewLegacyKeccak256()

	hashsymbol.Write([]byte("symbol()"))
	buf = hashsymbol.Sum(nil)

	callpara2 := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err = StormClient.EthCall(callpara2, "latest")
	if err != nil {
		return map[string]interface{}{}, err
	}
	symbol := new(string)
	returns = common.FromHex(tx)
	out, err = Abi.Unpack("symbol", returns)
	*symbol = fmt.Sprint(out[0]) // modified 2020/10/27
	return map[string]interface{}{"name": nameb, "symbol": *symbol, "supply": supply}, err
}

//send a 721 to other
func (s *Service) SendErc721Token(ctx context.Context, addr, pass, contract, tokenid, memo, targetid string, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, pass)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, targetid, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Println(from, " try to send token %s:%s to addr:%s", contract, tokenid, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)
	tid, err := strconv.ParseUint(tokenid, 10, 0)
	if err != nil {
		return "", err
	}
	if tid == 0 {
		return "", errors.New("tokenid is not correct!")
	}
	var m abi.Method
	var exists bool
	if len(memo) == 0 {
		m, exists = Abi.Methods["transferFrom"]
	} else {
		m, exists = Abi.Methods["transferWithMemo"]
	}

	if exists {
		fmt.Println("Function is expected", m)
	}

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(from, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(toAddress, &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(tokenid, &m.Inputs[2])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 2, m.Inputs[2].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	if len(memo) > 0 {
		v, err = paramDecode(memo, &m.Inputs[3])

		if err != nil {
			err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 3, m.Inputs[3].Name, err)
			return "", err
		}
		method_params = append(method_params, v)
	}

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC721TokenSend
	}
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(memo)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC721Limit * 2,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send 721 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Send toke return error")
		return "", ee
	}

}

// 获取 memo search by blocknumber
func (s *Service) GetSendErc721TokenMemo(ctx context.Context, contractaddr, tokenid, txhash string) (string, error) {

	numstr := IntStringToHex(tokenid)
	bufofnun := []byte(numstr)
	wc := "0x" + string(Padding(bufofnun[2:], 64))
	fmt.Println("tokenid is ", wc)
	hash := sha3.NewLegacyKeccak256()
	var buf []byte
	//hash.Write([]byte{0xcc})
	hash.Write([]byte("transactionMemo(uint256,string)"))
	buf = hash.Sum(nil)
	//fmt.Println(hex.EncodeToString(buf[:]))
	//fmt.Println(Sign721Log1)
	opcode := "0x" + hex.EncodeToString(buf[:])
	//tx, err := StormClient.EthGetLogs(filter)
	tx, err := StormClient.EthGetTransactionReceipt(txhash) //tx is filter id
	if err != nil {
		return "", err
	}
	logs := tx.Logs
	var retMemo string
	//fmt.Println(logs)
	for i := 0; i < len(logs); i++ {
		if logs[i].Topics[0] == opcode {
			dd := []byte(logs[i].Data)
			fmt.Println(dd, string(dd))
			if len(dd) < 128 {
				return "", fmt.Errorf("error in log data")
			}
			mm, _ := hex.DecodeString(string(dd[128:]))
			mm = UnPadding(mm)
			retMemo = fmt.Sprint(string(mm))
			break
		}
	}

	//mt, err := json.Marshal(logs)

	return retMemo, err
}

// 获取 memo list search by
func (s *Service) GetErc721TokenMemoList(ctx context.Context, contractaddr, tokenid string) (string, error) {
	//var buf [32]byte
	//wo := hex.EncodeToString([]byte("transactionMemo"))
	//wc := "0x" + string(Padding([]byte(wo), 64))

	numstr := IntStringToHex(tokenid)
	bufofnun := []byte(numstr)
	wc := "0x" + string(Padding(bufofnun[2:], 64))
	fmt.Println("tokenid is ", wc)
	hash := sha3.NewLegacyKeccak256()
	var buf []byte
	//hash.Write([]byte{0xcc})
	hash.Write([]byte("memoAdded(uint256,string)"))
	buf = hash.Sum(nil)
	fmt.Println(hex.EncodeToString(buf[:]))
	//fmt.Println(Sign721Log2)

	filter := storm.FilterParams{
		FromBlock: IntStringToHex("0"),
		ToBlock:   "latest",
		Address:   []string{contractaddr},
		Topics: []string{
			"0x" + hex.EncodeToString(buf[:]),
			wc,
		},
	}

	//tx, err := StormClient.EthGetLogs(filter)
	tx, err := StormClient.EthNewFilter(filter) //tx is filter id
	if err != nil {
		return "", err
	}
	logs, err := StormClient.EthGetFilterLogs(tx)

	for i := 0; i < len(logs); i++ {
		dd := []byte(logs[i].Data)

		mm, _ := hex.DecodeString(string(dd[128:]))
		mm = UnPadding(mm)
		logs[i].Data = fmt.Sprint(string(mm))
	}
	mt, err := json.Marshal(logs)
	//str := new(string)
	//err = Abi.Unpack(&str, "transactionMemo", common.FromHex(tx))
	//	mmemo := common.FromHex(tx)

	/* type log struct {
		TokenId big.Int
		Memo    string
	}

	retlog := make([]log, 0)
	for _, vLog := range logs {
		event := log{}

		fmt.Println(vLog.Data)
		err := Abi.Unpack(&event, "transactionMemo", []byte(vLog.Data))
		if err != nil {
			fmt.Println(err)
			continue
		}
		retlog = append(retlog, event)
		fmt.Println(event) // foo
	}
	mt, _ := json.Marshal(retlog) */
	return string(mt), err
}

// add memo to a 721 token
func (s *Service) AddErc721TokenMemo(ctx context.Context, userid, password, contract, tokenid, memo string, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, userid, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	fmt.Println(from, " try to set  new memo at ", contract, tokenid)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)
	tid, err := strconv.ParseUint(tokenid, 10, 0)
	if err != nil {
		return "", err
	}
	if tid == 0 {
		return "", errors.New("tokenid is not correct!")
	}
	m, exists := Abi.Methods["addMemo"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0]) //tokenid
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(memo, &m.Inputs[1]) //property

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC721TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(memo)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC721Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, userid)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", userid)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx mint  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}

		return t.TransactionHash, ee
	} else {
		return "", ee
	}

}

// 获取property
func (s *Service) GetErc721TokenProperty(ctx context.Context, contractaddr, tokenid string) (string, error) {

	m, exists := Abi.Methods["getProperty"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: contractaddr,
		To:   contractaddr,
		Data: hexutil.Encode(dataEncodefor721Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}
	nameb := new(string)
	returns := common.FromHex(tx)
	out, err := Abi.Unpack("getProperty", returns)
	*nameb = fmt.Sprint(out[0])
	return *nameb, err

}

// Get Erc721 token 's metadata
func (s *Service) GetErc721MetaData(ctx context.Context, contractaddr, tokenid string) (string, error) {

	m, exists := Abi.Methods["tokenURI"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: contractaddr,
		To:   contractaddr,
		Data: hexutil.Encode(dataEncodefor721Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}
	nameb := new(string)
	returns := common.FromHex(tx)
	out, err := Abi.Unpack("tokenURI", returns)
	*nameb = fmt.Sprint(out[0]) // modified 2020/10/27

	return *nameb, err

}
func (s *Service) GetTx(ctx context.Context, addr string) (*baas.BlockTx, error) {
	return &baas.BlockTx{
		Type:   "",
		Filter: "",
	}, nil
}

func (s *Service) GetBlockNumber(ctx context.Context) (string, error) {
	var err error

	nm, err := StormClient.EthBlockNumber()
	if err != nil {
		return "", errors.New("Get BlockNumber Error")
	}
	return fmt.Sprint(nm), nil

}

//
func (s *Service) GetUserErc721TokenList(ctx context.Context, userid, contractaddr string) (string, error) {
	Address, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, userid, "")
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	totalstr, err := s.GetErc721Balance(ctx, userid, contractaddr)
	if err != nil {
		return "", err
	}

	total, err := strconv.Atoi(totalstr)
	if err != nil {
		return "", err
	}
	returnData := make([]map[string]interface{}, 0)
	for i := 0; i < total; i++ {
		m, exists := Abi.Methods["tokenOfOwnerByIndex"]
		if exists {
			fmt.Println("Function is expected", m)
		}
		method_params := make([]interface{}, 0, 8)
		v, err := paramDecode(Address, &m.Inputs[0])
		if err != nil {
			err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
			return "", err
		}
		method_params = append(method_params, v)
		v, err = paramDecode(fmt.Sprint(i), &m.Inputs[1])
		if err != nil {
			err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
			return "", err
		}
		method_params = append(method_params, v)
		dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
		if err != nil {
			return "", err
		}

		callpara := storm.T{
			From: Address,
			To:   contractaddr,
			Data: hexutil.Encode(dataEncodefor721Call),
		}
		tx, err := StormClient.EthCall(callpara, "latest")
		if err != nil {
			return "", err
		}
		//	var nameb *big.Int
		returns := common.FromHex(tx)
		out, err := Abi.Unpack("tokenOfOwnerByIndex", returns)

		if err != nil {
			return "", err
		}
		returnData = append(returnData, map[string]interface{}{"tokenid": fmt.Sprint(out[0])}) // modified 2020/10/27

	}

	//	fmt.Println(*nameb)
	outputstr, err := json.Marshal(returnData)
	return string(outputstr), err
}

//Create new 721 token
func (s *Service) CreateERC721Contract(ctx context.Context, userid, pass, name, symbol, class string) (string, string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, userid, pass)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", "", err
	}

	s.Log.Debug().Msgf(" try to create new 721 contract with name: %s symbol:%s to addr:%s and type %s", userid, name, symbol, class)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", "", err
	}

	dataContract := getBaas721Code()

	if err != nil {
		return "", "", err
	}

	mm, err := Abi.Constructor.Inputs.PackValues([]interface{}{name, symbol})

	dataContract = append(dataContract, mm...)
	var maxGas int64

	maxGas = baas.DelplyContractGas
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		Gas:      int(maxGas), //3000000, //baas.GasERC721Limit * 10,
		GasPrice: big.NewInt(GasPrice),
		Value:    big.NewInt(0),
		Data:     string(dataContract),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, userid)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", userid)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to deploy  721")
			}
		}
		if t.Status == "0x0" {
			return "", t.TransactionHash, errors.New("Transaction failed")
		}
		return t.ContractAddress, t.TransactionHash, nil
	} else {
		s.Log.Err(err).Msg("Deploy 721 Contract  return error")
		return "", "", ee
	}

}

//user create new ERC20 contranct with type
func (s *Service) DeployERC20Contract(ctx context.Context, userid, password, name, symbol, class string, total uint64, decimal uint8, capacity uint64) (string, string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, userid, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", "", err
	}

	s.Log.Debug().Msgf(" try to create new 20 contract with name: %s symbol:%s to addr:%s and type %s,decimal %d,capacity:%d", userid, name, symbol, class, decimal, capacity)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", "", err
	}

	erc20Abi, err := abi.JSON(strings.NewReader(getBaas20ABIStr(class)))
	if err != nil {

		return "", "", fmt.Errorf("erc20 abi parse failed,please check")
	}

	dataContract := getBaas20Code(class)

	if err != nil {
		return "", "", err
	}

	var mm []byte
	if class == "4" {
		mm, err = erc20Abi.Constructor.Inputs.PackValues([]interface{}{name, symbol, decimal, big.NewInt(int64(total)), big.NewInt(int64(capacity))})
	} else {
		mm, err = erc20Abi.Constructor.Inputs.PackValues([]interface{}{name, symbol, decimal, big.NewInt(int64(total))})
	}

	dataContract = append(dataContract, mm...)
	var maxGas int64

	maxGas = baas.DelplyContractGas
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		Gas:      int(maxGas), //5000000,
		GasPrice: big.NewInt(GasPrice),
		Value:    big.NewInt(0),
		Data:     string(dataContract),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, userid)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", userid)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to deploy erc20")
			}
		}
		if t.Status == "0x0" {
			return "", t.TransactionHash, errors.New("Transaction failed")
		}
		return t.ContractAddress, t.TransactionHash, nil
	} else {
		s.Log.Err(err).Msg("Deploy 20 Contract  return error")
		return "", "", ee
	}
}

func (s *Service) GetErc20TotalSupply(ctx context.Context, addr string) (string, error) {

	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	//hash.Write([]byte{0xcc})
	hash.Write([]byte("totalSupply()"))
	buf = hash.Sum(nil)

	//fmt.Println(buf)

	callpara := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}

	//var balance *big.Int
	var returns = common.FromHex(tx)
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	out, err := erc20Abi.Unpack("totalSupply", returns)
	decimalstr, err := s.GetErc20Decimal(ctx, addr)

	return removeDecimal(fmt.Sprint(out[0]), decimalstr), err //modified 2020/10/27

}

func (s *Service) GetErc20Decimal(ctx context.Context, addr string) (string, error) {

	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	//hash.Write([]byte{0xcc})
	hash.Write([]byte("decimals()"))
	buf = hash.Sum(nil)

	//fmt.Println(buf)

	callpara := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	//var dd uint8
	var returns = common.FromHex(tx)
	out, err := erc20Abi.Unpack("decimals", returns)

	return fmt.Sprint(out[0]), err //modified 2020/10/27

}

func (s *Service) GetErc20Info(ctx context.Context, addr string) (map[string]interface{}, error) {
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return map[string]interface{}{}, err
	}
	supply, err := s.GetErc20TotalSupply(ctx, addr)
	if err != nil {
		return map[string]interface{}{}, err
	}
	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	hash.Write([]byte("name()"))
	buf = hash.Sum(nil)

	callpara := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return map[string]interface{}{}, err
	}

	nameb := new(string)
	var returns = common.FromHex(tx)
	out, err := erc20Abi.Unpack("name", returns)
	*nameb = fmt.Sprint(out[0]) // modified 2020/10/27
	hashsymbol := sha3.NewLegacyKeccak256()

	hashsymbol.Write([]byte("symbol()"))
	buf = hashsymbol.Sum(nil)

	callpara2 := storm.T{
		From: addr,
		To:   addr,
		Data: "0x" + hex.EncodeToString(buf[0:10]),
	}
	tx, err = StormClient.EthCall(callpara2, "latest")
	if err != nil {
		return map[string]interface{}{}, err
	}
	symbol := new(string)
	returns = common.FromHex(tx)
	out, err = erc20Abi.Unpack("symbol", returns)
	*symbol = fmt.Sprint(out[0])
	dd, err := s.GetErc20Decimal(ctx, addr)
	if err != nil {
		return map[string]interface{}{}, err
	}
	//decimal, _ := strconv.Atoi(dd)

	//amount, _ := big.NewFloat(0).SetString(supply)
	//deci := big.NewFloat(0).SetFloat64(math.Pow10(decimal))
	//realSupply := big.NewFloat(0).Quo(amount, deci)

	return map[string]interface{}{"name": nameb, "symbol": *symbol, "supply": supply, "decimals": dd}, err
}

// get user in erc20 contract hash
func (s *Service) GetErc20Balance(ctx context.Context, userid, addr string) (string, error) {
	address, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, userid, "")
	if err != nil {
		s.Log.Err(err).Msg("From User id is not found")
		return "", err
	}
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	m, exists := erc20Abi.Methods["balanceOf"]
	if exists {
		fmt.Println("Function is expected", m)
	}

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(address, &m.Inputs[0])
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: address,
		To:   addr,
		Data: hexutil.Encode(dataEncodefor721Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}

	//var balance *big.Int
	var returns = common.FromHex(tx)
	out, err := erc20Abi.Unpack("balanceOf", returns)

	decimalstr, err := s.GetErc20Decimal(ctx, addr)
	decimal, _ := strconv.Atoi(decimalstr)

	//f := new(big.Float).SetInt(balance)
	f, _ := new(big.Float).SetString(fmt.Sprint(out[0])) // modified 2020/10/27

	d := big.NewFloat(math.Pow10(decimal))
	return f.Quo(f, d).String(), err
}

// send erc20 token to others

func (s *Service) SendErc20Token(ctx context.Context, addr, password, toAddr, contract, memo string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Printf("%s try to send erc20 token %s to addr:%s", from, contract, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	var exists bool
	if len(memo) == 0 {
		m, exists = erc20Abi.Methods["transfer"]
	} else {
		m, exists = erc20Abi.Methods["transferWithMemo"]
	}

	if exists {
		fmt.Println("Function is expected", m)
	}
	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)
	v1, v2 := big.NewFloat(math.Pow10(decimal)), big.NewFloat(value)
	v3 := v1.Mul(v1, v2)
	v3Str := v3.Text('f', 0)
	vinput, _ := big.NewInt(0).SetString(v3Str, 10) //(int64(value * math.Pow10(decimal)))

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(toAddress, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(vinput.String(), &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	if len(memo) > 0 {
		v, err = paramDecode(memo, &m.Inputs[2])
		if err != nil {
			err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 3, m.Inputs[2].Name, err)
			return "", err
		}
		method_params = append(method_params, v)
	}

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenSend
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	//	maxGas += int64(len(memo)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasMsgLimit * 2,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Send toke return error")
		return "", ee
	}

}

// ERC20 Approve interface
func (s *Service) ApproveErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Printf("%s try to send erc20 token %s to addr:%s", from, contract, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	var exists bool

	m, exists = erc20Abi.Methods["approve"]

	if exists {
		fmt.Println("Function is expected", m)
	}
	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)
	v1, v2 := big.NewFloat(math.Pow10(decimal)), big.NewFloat(value)
	v3 := v1.Mul(v1, v2)
	v3str := v3.Text('f', 0)
	fmt.Println(v3str)
	vinput, _ := big.NewInt(0).SetString(v3str, 10)

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(toAddress, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(vinput.String(), &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)
	//dataEncodeFor20Call, err := m.Inputs.PackValues([]interface{}{toAddress, vinput.String()})
	if err != nil {
		return "", err
	}
	//fmt.Println(dataEncodefor721Call)
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* 	if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), // baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Send toke return error")
		return "", ee
	}

}

func (s *Service) AllowanceErc20(ctx context.Context, addr, toAddr, contract string) (string, error) {
	from, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, addr, "")
	if err != nil {
		s.Log.Err(err).Msg("Send User id is not found")
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}

	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	m, exists := erc20Abi.Methods["allowance"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(from, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(toAddress, &m.Inputs[1])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor20Call, err := erc20Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: toAddress,
		To:   contract,
		Data: hexutil.Encode(dataEncodefor20Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}
	//var amount *big.Int
	var returns = common.FromHex(tx)

	//nameb := new(string)
	//	returns := common.FromHex(tx)
	decimalstr, _ := s.GetErc20Decimal(ctx, contract)

	out, err := erc20Abi.Unpack("allowance", returns)

	return removeDecimal(fmt.Sprint(out[0]), decimalstr), err

}

// ERC20 Approve interface
func (s *Service) IncreaseAllowanceErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Printf("%s try to send erc20 token %s to addr:%s", from, contract, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	var exists bool

	m, exists = erc20Abi.Methods["increaseAllowance"]

	if !exists {
		return "", fmt.Errorf("contract method not exists")
	} else {
		fmt.Println("Function is expected", m)
	}
	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)
	v1 := big.NewFloat(math.Pow10(decimal))
	v2 := big.NewFloat(0).SetFloat64(value)
	v3 := v1.Mul(v1, v2)
	v4 := v3.Text('f', 0)
	vinput, _ := big.NewInt(0).SetString(v4, 10)

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(toAddress, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(vinput.String(), &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)
	//dataEncodeFor20Call, err := m.Inputs.PackValues([]interface{}{toAddress, vinput.String()})
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Send toke return error")
		return "", ee
	}

}

// ERC20 Approve interface
func (s *Service) DecresaseAllowanceErc20(ctx context.Context, addr, password, toAddr, contract string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Printf("%s try to send erc20 token %s to addr:%s", from, contract, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	var exists bool

	m, exists = erc20Abi.Methods["decreaseAllowance"]

	if exists {
		fmt.Println("Function is expected", m)
	}
	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)

	v1 := big.NewFloat(math.Pow10(decimal))
	v2 := big.NewFloat(0).SetFloat64(value)
	v3 := v1.Mul(v1, v2)
	v4 := v3.Text('f', 0)
	vinput, _ := big.NewInt(0).SetString(v4, 10)
	//vinput := big.NewInt(int64(value * math.Pow10(decimal)))

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(toAddress, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(vinput.String(), &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)
	//dataEncodeFor20Call, err := m.Inputs.PackValues([]interface{}{toAddress, vinput.String()})
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Send toke return error")
		return "", ee
	}

}

func (s *Service) TransferFromErc20(ctx context.Context, user, password, fromaddr, toAddr, contract, memo string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, user, password)
	//fmt.Println(frompk)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}
	//fmt.Println(frompk)
	fromAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, fromaddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}

	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, toAddr, "")
	if err != nil {
		s.Log.Err(err).Msg("Target User id is not found")
		return "", err
	}
	fmt.Printf("%s try to transfer erc20 token %s from %s to addr:%s", from, contract, fromAddress, toAddress)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("1")
	if err != nil {
		return "", err
	}
	var exists bool
	if len(memo) == 0 {
		m, exists = erc20Abi.Methods["transferFrom"]
	} else {
		m, exists = erc20Abi.Methods["transferFromWithMemo"]
	}

	if exists {
		fmt.Println("Function is expected", m)
	}

	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)

	v1 := big.NewFloat(math.Pow10(decimal))
	v2 := big.NewFloat(0).SetFloat64(value)
	v3 := v1.Mul(v1, v2)
	v4 := v3.Text('f', 0)
	vinput, _ := big.NewInt(0).SetString(v4, 10)
	var dataEncodeforCall []byte
	fmt.Println(m.RawName)
	/* if len(memo) == 0 {
		dataEncodeforCall, err = m.Inputs.PackValues([]interface{}{common.HexToAddress(fromAddress), common.HexToAddress(toAddress), vinput})
	} else {
		dataEncodeforCall, err = m.Inputs.PackValues([]interface{}{common.HexToAddress(fromAddress), common.HexToAddress(toAddress), vinput, memo})
	}

	if err != nil {
		return "", err
	} */
	//vinput := big.NewInt(int64(value * math.Pow10(decimal)))

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(fromAddress, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(toAddress, &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	v, err = paramDecode(vinput.String(), &m.Inputs[2])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 2, m.Inputs[2].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	if len(memo) > 0 {
		v, err = paramDecode(memo, &m.Inputs[3])
		if err != nil {
			err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 3, m.Inputs[3].Name, err)
			return "", err
		}
		method_params = append(method_params, v)
	}

	dataEncodeforCall, err = erc20Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenSend
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	//maxGas += int64(len(memo)) * 12
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasMax,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodeforCall),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, user)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", err)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("TransferFrom toke return error")
		return "", ee
	}
}

func (s *Service) GetErc20TxMemo(ctx context.Context, hash string) (string, error) {
	tx, err := StormClient.EthGetTransactionReceipt(hash) //tx is filter id
	if err != nil {
		return "", err
	}
	logs := tx.Logs
	var retMemo string
	//fmt.Println(logs)
	for i := 0; i < len(logs); i++ {
		if logs[i].Topics[0] == "0x35fad523fac4c32a939871d91c2d88ee509d2f04fdb4e41ad07d8bb7b96e490d" {
			dd := []byte(logs[i].Data)
			fmt.Println(string(dd))
			if len(dd) < 64 {
				return "", fmt.Errorf("no memo")
			}
			mm, _ := hex.DecodeString(string(dd[128:]))
			mm = UnPadding(mm)
			retMemo = fmt.Sprint(string(mm))
			break
		}
	}

	return retMemo, err

}

func (s *Service) BurnErc20(ctx context.Context, addr, password, contract string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}

	fmt.Printf("%s try to Burn erc20 token %s to addr:%s", from, contract)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("2")
	if err != nil {
		return "", err
	}
	var exists bool

	m, exists = erc20Abi.Methods["burn"]
	if exists {
		fmt.Println("Function is expected", m)
	} else {
		return "", fmt.Errorf("contract method is not support")
	}
	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)
	v1, v2 := big.NewFloat(math.Pow10(decimal)), big.NewFloat(value)
	v3 := v1.Mul(v1, v2)
	v3str := v3.Text('f', 0)
	//	fmt.Println(v3str)
	vinput, _ := big.NewInt(0).SetString(v3str, 10)

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(vinput.String(), &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)

	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx burn erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("burn token return error")
		return "", ee
	}
}

func (s *Service) PauseErc20(ctx context.Context, addr, password, contract string, value bool, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}

	fmt.Printf("%s try to pause erc20 token %s to addr:%s status:%v", from, contract, value)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("2")
	if err != nil {
		return "", err
	}
	var exists bool
	if value {
		m, exists = erc20Abi.Methods["pause"]
	} else {
		m, exists = erc20Abi.Methods["unpause"]
	}

	if exists {
		fmt.Println("Function is expected", m)
	} else {
		return "", fmt.Errorf("contract method is not support")
	}

	method_params := make([]interface{}, 0, 8)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)

	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx puase erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("puase token return error")
		return "", ee
	}
}

func (s *Service) GetErc20PauseStatus(ctx context.Context, contract string) (string, error) {

	erc20Abi, err := getBaas20ABI("2")
	if err != nil {
		return "", err
	}
	m, exists := erc20Abi.Methods["paused"]
	if exists {
		fmt.Println("Function is expected", m)
	} else {
		return "", fmt.Errorf("contract method is not support")
	}

	method_params := make([]interface{}, 0, 8)

	dataEncodefor20Call, err := erc20Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}

	callpara := storm.T{
		From: contract,
		To:   contract,
		Data: hexutil.Encode(dataEncodefor20Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}
	//var statusOfContract bool
	var returns = common.FromHex(tx)

	out, err := erc20Abi.Unpack("paused", returns)

	return fmt.Sprint(out[0]), err //modified 2020/10/27

}

func (s *Service) MintErc20(ctx context.Context, addr, password, contract string, value float64, gas int64) (string, error) {
	var err error
	from, frompk, rv, cipher, salt, err := s.BlockStorage.GetAddressByService(ctx, addr, password)
	if err != nil {
		s.Log.Err(err).Msg("User id is not found")
		return "", err
	}
	frompk, err = s.DecryptKey(frompk, "", cipher, rv, salt)
	if err != nil {
		return "", err
	}

	fmt.Printf("%s try to mint erc20 token %s to addr:%s quantity:%f", from, contract, value)

	balancestr, err := s.GetBalance(ctx, from)
	f, _ := strconv.ParseFloat(balancestr, 64)
	if GasPrice > 0 && f < GasCheckMinimum { //TODO:need check mint token total gas
		return "", baas.ErrBaasNotEnoughMoney
	}

	nounce, err := StormClient.EthGetTransactionCount(from, "latest")
	if err != nil {
		return "", err
	}

	vv, err := strconv.ParseFloat("0", 64)
	amount := FloatToBigInt(vv)

	var m abi.Method
	erc20Abi, err := getBaas20ABI("3")
	if err != nil {
		return "", err
	}
	var exists bool

	m, exists = erc20Abi.Methods["mint"]
	if exists {
		fmt.Println("Function is expected", m)
	} else {
		return "", fmt.Errorf("contract method is not support")
	}

	decimalstr, err := s.GetErc20Decimal(ctx, contract)
	decimal, _ := strconv.Atoi(decimalstr)
	v1, v2 := big.NewFloat(math.Pow10(decimal)), big.NewFloat(value)
	v3 := v1.Mul(v1, v2)
	v3str := v3.Text('f', 0)
	//fmt.Println(v3str)
	vinput, _ := big.NewInt(0).SetString(v3str, 10)

	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(from, &m.Inputs[0])
	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err)
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(vinput.String(), &m.Inputs[1])

	if err != nil {
		err := fmt.Errorf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err)
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := erc20Abi.Pack(m.Name, method_params...)
	//dataEncodeFor20Call, err := m.Inputs.PackValues([]interface{}{toAddress, vinput.String()})
	if err != nil {
		return "", err
	}
	var maxGas = gas
	if maxGas <= 0 {
		maxGas = baas.ERC20TokenOP
	}
	/* if GasPrice == 0 {
		maxGas = 0
	} */
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      int(maxGas), //baas.GasERC20Limit,
		GasPrice: big.NewInt(GasPrice),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     string(dataEncodefor721Call),
		Nonce:    nounce,
	}

	t, err := SignAndSendTx(tx, frompk)
	var ee = err
	if ee == nil {
		if len(t.TransactionHash) == 0 {
			ee = errors.New("transaction timeout")
		}
		if ee == nil {
			appid, err := s.BlockStorage.GetApplicationId(ctx, addr)
			if err != nil {
				fmt.Println("Warnning, Application id search error:", addr)
			}
			updatedUsages := &baas.UpdateUsage{
				ApplicationID: appid,
				TxCount:       1,
			}

			if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
				s.Log.Error().Err(err).Msg("failed to update block tx send erc20 token  usage")
			}
		}
		if t.Status == "0x0" {
			return t.TransactionHash, errors.New("Transaction failed")
		}
		return t.TransactionHash, ee
	} else {
		s.Log.Err(err).Msg("Mint token return error")
		return "", ee
	}
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TxList(ctx context.Context, contract, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndPeerToERC721Tx, "*", contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend get erc 721 tx url is  ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

// encaspulate tx to json string and return

func (s *Service) GetErc721TokenTxList(ctx context.Context, contract, tokenid, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndPeerERC721TokenTx, "*", tokenid, contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend url for 721 token tx query: ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TxListByUser(ctx context.Context, userid, contract, page, size string) (string, error) {
	var err error
	address := userid

	if !strings.HasPrefix(address, "0x") {
		address, _, _, _, _, err = s.BlockStorage.GetAddressByService(ctx, userid, "")
		if err != nil {
			s.Log.Err(err).Msg("From User id is not found")
			return "", err
		}
	}

	url := fmt.Sprintf(BackEndPeerToERC721Tx, address, contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Println("backend service for 721 with user addres  url is : ", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TokenTxListByUser(ctx context.Context, userid, contract, tokenid, page, size string) (string, error) {
	var err error
	toAddress, _, _, _, _, err := s.BlockStorage.GetAddressByService(ctx, userid, "")
	if err != nil {
		s.Log.Err(err).Msg("request User id is not found")
		return "", err
	}
	//url := fmt.Sprintf(BackEndGetAddressTx, addr)
	url := fmt.Sprintf(BackEndPeerERC721TokenTx, toAddress, tokenid, contract)
	url += fmt.Sprintf("?page=%s", page)
	fmt.Printf("backend url to get token tx for specified user  %s", url)
	resp, err := http.Get(url)

	if err != nil {
		s.Log.Err(err).Msg("Backend system error!")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		s.Log.Err(err).Msg("Backend system return data error!")
		return "", err
	}

	code, ok := result["statusCode"].(string)
	if ok {
		if code == "1" {
			dataArray := result["data"]
			datastr, err := json.Marshal(dataArray)
			if err != nil {
				s.Log.Err(err).Msg("Backend system return data error!")
				return "", err
			}
			return string(datastr), nil

		}
	}
	return "", baas.ErrBaasQueryNoResult
}

func (s *Service) GetApplicationUsers(ctx context.Context, id string, page, limit int) ([]*baas.BAASUser, error) {
	if limit == 0 {
		limit = baas.DefaultFilePaginationLimit
	}

	return s.BlockStorage.ApplicationUsers(ctx, id, page, limit)
}

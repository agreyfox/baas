package block

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agreyfox/baas"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"

	"github.com/agreyfox/baas/cmd/baasd"
	"github.com/agreyfox/baas/storm"
	"github.com/rs/zerolog"
)

var (
	StormClient *storm.EthRPC
	BigHelper   = big.NewInt(10)
	Wei         = BigHelper.Exp(BigHelper, big.NewInt(18), nil)
	Abi         abi.ABI
)

const (
	BackEndGetAddressTx = "http://dao.moacchain.net/api/v1/wallets/%s/tx"
	BackEndPeerToPeerTx = "http://dao.moacchain.net/api/v1/wallets/%s/%s/txfilter"

	ContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
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
	Abi, err = abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		fmt.Println("721 abi parse failed,please check ")
		return
	}
	fmt.Println(Abi)
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

func (s *Service) SendToken(ctx context.Context, addr, password, toAddr, value, msg string) (string, error) {
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
	if len(msg) > 0 {
		datappayload = msg
		datalen = len(datappayload)
	} else {
		datappayload = ""
		datalen = 0
	}
	//fmt.Println(storm.BigToHex(vv))
	tx := storm.T{
		From:     from,
		To:       to,
		Gas:      baas.GasLimit + 3000 + datalen*6,
		GasPrice: big.NewInt(5000000000),
		Value:    amount, //big.NewInt(1000000000000000000),
		Data:     datappayload,
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
		return t.TransactionHash, ee
	} else {
		return "", ee
	}
	// blocknm, err := StormClient.EthSendTransaction(tx)

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
		return t.TransactionHash, ee
	} else {
		return "", ee
	}

}

/// read msg from blockchain
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
	return "", errors.New("No result return")
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
	return "", errors.New("No result return")
}

func (s *Service) GetErc20Balance(ctx context.Context, addr string) (string, error) { return "", nil }

func (s *Service) SendErc20Token(ctx context.Context, addr string) (string, error)   { return "", nil }
func (s *Service) CreateErc20Token(ctx context.Context, addr string) (string, error) { return "", nil }
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
	//fmt.Println(dataEncodefor721Call)
	//fmt.Println(hexutil.Encode(dataEncodefor721Call))
	callpara := storm.T{
		From: address,
		To:   contractAddr,
		Data: hexutil.Encode(dataEncodefor721Call),
	}
	tx, err := StormClient.EthCall(callpara, "latest")
	if err != nil {
		return "", err
	}

	var balance *big.Int
	var returns = common.FromHex(tx)
	err = Abi.Unpack(&balance, "balanceOf", returns)
	//fmt.Println(balance)
	return balance.String(), err

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

	var balance *big.Int
	var returns = common.FromHex(tx)
	err = Abi.Unpack(&balance, "totalSupply", returns)
	//fmt.Println(balance)
	return balance.String(), err
	//	return "", nil
}

func (s *Service) CreateErc721Token(ctx context.Context, userid, password, contract, tokenid, meta string) (string, error) {
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
	if f < 0.001 { //TODO:need check mint token total gas
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
	v, err := paramDecode(tokenid, &m.Inputs[0])
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(meta, &m.Inputs[1])

	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err))
		return "", err
	}
	method_params = append(method_params, v)

	dataEncodefor721Call, err := Abi.Pack(m.Name, method_params...)
	if err != nil {
		return "", err
	}
	//fmt.Println(dataEncodefor721Call)
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      baas.GasERC721Limit,
		GasPrice: big.NewInt(5000000000),
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
		return t.TransactionHash, ee
	} else {
		return "", ee
	}
	return "", nil
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
	err = Abi.Unpack(&nameb, "name", returns)

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
	Abi.Unpack(&symbol, "symbol", returns)
	return map[string]interface{}{"name": nameb, "symbol": *symbol, "supply": supply}, err
}

//send a 721 to other
func (s *Service) SendErc721Token(ctx context.Context, addr, pass, contract, tokenid, targetid string) (string, error) {
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
	if f < 0.001 { //TODO:need check mint token total gas
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
	m, exists := Abi.Methods["transferFrom"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(from, &m.Inputs[0])
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 0, m.Inputs[0].Name, err))
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(toAddress, &m.Inputs[1])

	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode parameter %v (%v): %v", 1, m.Inputs[1].Name, err))
		return "", err
	}
	method_params = append(method_params, v)
	v, err = paramDecode(tokenid, &m.Inputs[2])

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
	tx := storm.T{
		From:     from,
		To:       contract,
		Gas:      baas.GasERC721Limit,
		GasPrice: big.NewInt(5000000000),
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
		return t.TransactionHash, ee
	} else {
		return "", ee
	}

}
func (s *Service) GetErc721MetaData(ctx context.Context, contractaddr, tokenid string) (string, error) {

	m, exists := Abi.Methods["tokenURI"]
	if exists {
		fmt.Println("Function is expected", m)
	}
	method_params := make([]interface{}, 0, 8)
	v, err := paramDecode(tokenid, &m.Inputs[0])
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
	err = Abi.Unpack(&nameb, "tokenURI", returns)

	//	fmt.Println(*nameb)
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

// encaspulate tx to json string and return
func (s *Service) GetErc721TxList(ctx context.Context, addr, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndGetAddressTx, addr)
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
	return "", errors.New("No result return")
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TokenTxList(ctx context.Context, addr, tokenid, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndGetAddressTx, addr)
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
	return "", errors.New("No result return")
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TxListByUser(ctx context.Context, addr, userid, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndGetAddressTx, addr)
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
	return "", errors.New("No result return")
}

// encaspulate tx to json string and return
func (s *Service) GetErc721TokenTxListByUser(ctx context.Context, addr, tokenid, userid, page, size string) (string, error) {
	var err error

	url := fmt.Sprintf(BackEndGetAddressTx, addr)
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
	return "", errors.New("No result return")
}

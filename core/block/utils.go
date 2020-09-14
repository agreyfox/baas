package block

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/agreyfox/baas"
	"github.com/agreyfox/baas/storm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	NetworkID = "4452999"
)

// sign struct for raw message
func SignAndSendTx(st storm.T, privateKey string) (*storm.TransactionReceipt, error) {

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Errorf("%s", err)
		return nil, err
	}

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}

	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fromAddress := st.From
	nonce := uint64(st.Nonce)
	/* nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	*/
	value := st.Value //big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(baas.GasMsgLimit)
	gasPrice := st.GasPrice // in units
	/* gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	} */
	toAddress := common.HexToAddress(st.To)
	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	data = []byte(st.Data)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, _ := storm.ParseBigInt(NetworkID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(&chainID), pk)
	if err != nil {
		fmt.Printf("Transaction error:%s", err)
		return nil, err
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := "0x" + hex.EncodeToString(rawTxBytes)

	txhash, err := StormClient.EthSendRawTransaction(rawTxHex) /// SendRawTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Printf("Transaction error:%s", err)
		return nil, err
	}
	//fmt.Println(txhash)
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	re, err := StormClient.EthGetTransactionReceipt(txhash)
	var timeout int
	for len(re.TransactionHash) == 0 && timeout <= 30 {
		re, _ = StormClient.EthGetTransactionReceipt(txhash)
		time.Sleep(time.Duration(2) * time.Second)
		timeout += 1
		fmt.Printf("%d", timeout)
	}
	//fmt.Println(re, err)
	return re, nil
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func paramDecode(param string, arg *abi.Argument) (v interface{}, err error) {
	param = strings.TrimSpace(param)
	fmt.Println(arg.Type)
	switch arg.Type.T {
	case abi.StringTy:
		str_val := new(string)
		v = str_val
		//err = json.Unmarshal([]byte(param), &v)
		v = fmt.Sprintf(param)
		err = nil
	case abi.UintTy, abi.IntTy:
		val := big.NewInt(0)
		_, success := val.SetString(param, 10)
		if !success {
			err = errors.New(fmt.Sprintf("Invalid numeric (base 10) value: %v", param))
		}
		v = val
	case abi.AddressTy:
		if !((len(param) == (common.AddressLength*2 + 2)) || (len(param) == common.AddressLength*2)) {
			err = errors.New(fmt.Sprintf("Invalid address length (%v), must be 40 (unprefixed) or 42 (prefixed) chars", len(param)))
		} else {
			var addr common.Address
			if len(param) == (common.AddressLength*2 + 2) {
				addr = common.HexToAddress(param)
			} else {
				var data []byte
				data, err = hex.DecodeString(param)
				addr.SetBytes(data)
			}
			v = addr
		}
	case abi.HashTy:
		if !((len(param) == (common.HashLength*2 + 2)) || (len(param) == common.HashLength*2)) {
			err = errors.New(fmt.Sprintf("Invalid hash length, must be 64 (unprefixed) or 66 (prefixed) chars"))
		} else {
			var hash common.Hash
			if len(param) == (common.HashLength*2 + 2) {
				hash = common.HexToHash(param)
			} else {
				var data []byte
				data, err = hex.DecodeString(param)
				hash.SetBytes(data)
			}
			v = hash
		}
	case abi.BytesTy:
		if len(param) > 2 {
			if (param[0] == '0') && (param[1] == 'x') {
				param = param[2:] // cut 0x prefix
			}
		}
		decoded_bytes, tmperr := hex.DecodeString(param)
		v = decoded_bytes
		err = tmperr
	case abi.BoolTy:
		val := new(bool)
		v = val
		err = json.Unmarshal([]byte(param), v)
	default:
		err = errors.New(fmt.Sprintf("Not supported parameter type: %v", arg.Type))
	}
	return v, err
}

func Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte('0')}, padding)
	return append(padtext, src...)
}

func UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	fsrc := src
	if src[0] == '0' && src[1] == 'x' {
		fsrc = src[2:]
	}
	length := len(fsrc)
	start := 0
	for i := 0; i < length; i++ {
		if fsrc[i] == 0 {
			start++
		} else {
			break
		}
	}

	end := 0
	for j := 1; j < length; j++ {

		if fsrc[length-j] == 0 {
			end++
		} else {
			break
		}
	}
	fmt.Print(fsrc[start : length-end])
	return fsrc[start : length-end]

}

func IntStringToHex(str string) string {
	//var n int64
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error when convert int string to hex")
		return ""
	}
	//return []byte(strconv.FormatInt(n, 16))
	return fmt.Sprintf("%#x", n)
}

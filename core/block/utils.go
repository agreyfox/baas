package block

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/agreyfox/baas"
	"github.com/agreyfox/baas/storm"
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
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := "0x" + hex.EncodeToString(rawTxBytes)

	txhash, err := StormClient.EthSendRawTransaction(rawTxHex) /// SendRawTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
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

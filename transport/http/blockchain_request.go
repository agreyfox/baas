package http

import (
	"errors"

	"github.com/agreyfox/baas"
	"github.com/fatih/color"
)

type createBaasUserRequest struct {
	Email            string `json:"userId"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	CipherText       string `json:"cipherText"`
	ApplicationID    string `json:"application_id"`
	Description      string `json:"description"`
	StorageAccessKey string `json:"storageAccessKey" `
	StorageSecretKey string `json:"storageSecretKey"`
	StorageRegion    string `json:"storageRegion" `
	StorageEngine    string `json:"storageEngine"`
	DeliveryURL      string `json:"deliveryUrl"`
}

func (r *createBaasUserRequest) validate() error {
	if r.Email == "" {
		return errors.New("Email is required")
	}
	if r.Password == "" || len(r.Password) < 8 {
		return errors.New("Password is required or not secure")
	}
	if r.ApplicationID == "" {
		return baas.ErrBaasApplicationIDRequired
	}
	if r.CipherText == "" || len(r.CipherText) < 8 {
		color.Red("No CiperText, Will use system key")
		return baas.ErrBaasCipherTextRequired
	}
	return nil
}

type updateBaasUserRequest struct {
	ID            string `json:"id"`
	Email         string `json:"userId"`
	OldPassword   string `json:"oldPassword"`
	NewPassword   string `json:"newPassword"`
	ApplicationID string `json:"application_id"`
	Password      string `json:"password"`
}

func (r *updateBaasUserRequest) validate() error {
	if r.ApplicationID == "" {
		return errors.New("application_id is required")
	}

	if r.OldPassword == "" || r.NewPassword == "" {
		return errors.New("password can't be empty")
	}

	if r.Email == "" {
		return errors.New("userId is required")
	}

	return nil
}

/*
type listApplicationQueryParam struct {
	Limit   int    `schema:"limit"`
	SinceID string `schema:"since_id"`
}
*/

type BlockOperation struct {
	Opr        string  `json:"op"`
	Userid     string  `json:"userId"`
	Password   string  `json:"password"`
	Targetid   string  `json:"toUserId"`
	CipherText string  `json:"cipherText"`
	Message    string  `json:"msg"`
	Encode     string  `json:"encode"`
	Value      float64 `json:"quantity"`
	Hash       string  `json:"txHash"`
	Size       string  `json:"total"`
	Page       string  `json:"page"`
	Contract   string  `json:"contract"`
	TokenId    string  `json:"tokenId"`
	GasLimit   int64   `json:"gasLimit"`
}

type SmartContractOperation struct {
	Contract    string  `json:"contract"`
	Meta        string  `json:"metadata"`
	UserId      string  `json:"userId"`
	Password    string  `json:"password"`
	TokenId     string  `json:"tokenId"`
	SendUserId  string  `json:"sendUserId"`
	TargetId    string  `json:"toUserId"`
	Property    string  `json:"property"`
	Memo        string  `json:"memo"`
	BlockNumber string  `json:"blockNumber"`
	Hash        string  `json:"txHash"`
	Quantity    float64 `json:"quantity"`
	Action      string  `json:"action"`
	GasLimit    int64   `json:"gasLimit"`
}

type DeployOperation struct {
	Name        string `json:"name"`
	UserId      string `json:"userId"`
	Password    string `json:"password"`
	Symbol      string `json:"symbol"`
	Decimal     uint8  `json:"decimals"`
	TotalSupply uint64 `json:"totalSupply"`
	Capacity    uint64 `json:"cap"`
	Class       uint8  `json:"type"`
	Description string `json:"desc"`
}

type TokenInfo struct {
	Status      string `json:"status"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply int    `json:"totalSupply"`
}

type BlockChainUsageQuery struct {
	ApplicationID string `json:"application_id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

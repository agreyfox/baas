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
	if r.Password == "" || len(r.Password) < 6 {
		return errors.New("Password is required or not secure")
	}
	if r.ApplicationID == "" {
		return baas.ErrBaasApplicationIDRequired
	}
	if r.CipherText == "" {
		color.Red("No CiperText, Will use system key")
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
		return errors.New("id is required")
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
	Opr        string `json:"op"`
	Userid     string `json:"userId"`
	Password   string `json:"password"`
	Targetid   string `json:"toUserId"`
	CipherText string `json:"cipherText"`
	Message    string `json:"message"`
	Value      string `json:"quantity"`
	Hash       string `json:"txHash"`
	Size       int    `json:"total"`
	Page       int    `json:"page"`
}

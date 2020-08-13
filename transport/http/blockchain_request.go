package http

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type createBaasUserRequest struct {
	Email            string `json:"userid"`
	Name             string `json:"name" `
	Password         string `json:"password"`
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
		return errors.New("Application is required")
	}

	return nil
}

type updateBaasUserRequest struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name" `
	Password    string `json:"password"`
	Description string `json:"description"`
}

func (r *updateBaasUserRequest) validate() error {
	if r.ID == "" {
		return errors.New("id is required")
	}

	if !govalidator.IsUUIDv4(r.ID) {
		return errors.New("invalid id")
	}

	if r.Name == "" {
		return errors.New("name is required")
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
	Opr      string `json:"op"`
	Userid   string `json:"userId"`
	Targetid string `json:"toUserId"`
	Message  string `json:"message"`
	Value    string `json:"quantity"`
	Hash     string `json:"hash"`
	Size     int    `json:"total"`
	Page     int    `json:"page"`
}

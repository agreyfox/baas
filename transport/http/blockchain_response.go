package http

import (
	"time"

	"github.com/agreyfox/baas"
)

type baasUserResponse struct {
	Status int           `json:"status"`
	Data   *bassUserData `json:"data"`
}

type baasUsersResponse struct {
	PaginationData PaginationData  `json:"meta"`
	Data           []*bassUserData `json:"data"`
}

type bassUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	ApplicationID string `json:"application_id"`
	Description   string `json:"description"`
	//StorageAccessKey string    `json:"storageAccessKey"`
	//StorageBucket    string    `json:"storageBucket"`
	//StorageEndpoint  string    `json:"storageEndpoint"`
	//StorageRegion    string    `json:"storageRegion"`
	//StorageEngine    string    `json:"storageEngine"`
	//DeliveryURL      string    `json:"cdnUrl"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func sanitizeBaasUser(v *baas.BAASUser) *bassUserData {
	return &bassUserData{
		ID:            v.ID,
		Name:          v.Name,
		Description:   v.Description,
		Email:         v.Email,
		ApplicationID: v.ApplicationID,
		Address:       v.Address,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
	}
}

func sanitizeBaasUsers(v []*baas.BAASUser) []*bassUserData {
	ss := make([]*bassUserData, len(v))
	for x, u := range v {
		nu := sanitizeBaasUser(u)
		ss[x] = nu
	}

	return ss
}

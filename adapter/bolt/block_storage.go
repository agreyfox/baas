package bolt

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/agreyfox/baas"
	"github.com/asdine/storm/v3/q"

	uuid "github.com/satori/go.uuid"

	"github.com/asdine/storm/v3"
	"golang.org/x/crypto/bcrypt"
)

var (
	r = mrand.New(mrand.NewSource(time.Now().Unix()))
)

type BlockStorage struct {
	DB *storm.DB
}

type blockUser struct {
	Pk            int    `storm:"id,increment"`
	ID            string `storm:"index"`
	ApplicationID string
	FileID        string
	Address       string //block address
	PrivateKey    string //block encrypt pk
	UserName      string `storm:"index"`
	Email         string `storm:"unique"`
	Salt          string //block encrypt salt
	CipherText    string //block encrypt cipherText
	Password      string // user password
	Rv            string //helper for encrypt algorithm
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t blockUser) validate(db *storm.DB) error {
	if t.ApplicationID == "" {
		return errors.New("applicationID is required")
	}

	if t.Password == "" {
		return errors.New("Password is required")
	}

	if t.Email == "" {
		return errors.New("Email is required")
	}

	var tran blockUser
	if err := db.Select(q.Eq("Email", t.Email)).First(&tran); err != nil {
		if err == storm.ErrNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (s *BlockStorage) Store(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	/* 	salt, err := randSalt()
	   	if err != nil {
	   		return nil, err
	   	}
	*/
	salt, _ := base64.StdEncoding.DecodeString(n.Salt)
	hash, err := hashPassword([]byte(n.Password), salt)
	if err != nil {
		return nil, err
	}

	t := blockUser{
		ID:            uuid.NewV4().String(),
		ApplicationID: n.ApplicationID,
		PrivateKey:    n.PrivateKey,
		Email:         n.Email,
		Address:       n.Address,
		Password:      string(hash),
		Salt:          n.Salt,
		Rv:            n.Rv,
		CipherText:    n.CipherText,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := t.validate(s.DB); err != nil {
		return nil, err
	}

	if err := s.DB.Save(&t); err != nil {
		if err == storm.ErrAlreadyExists {
			return nil, baas.ErrBaasUserNameTaken
		}
		return nil, err
	}

	mahiTran := sanitizeblockUser(t)

	return &mahiTran, nil
}

func (s *BlockStorage) UpdatePassword(ctx context.Context, id, old, newpass string) error {
	fmt.Printf("User try to Change password,", id)
	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return baas.ErrBaasNoSuchUser
		}
		return err
	}
	salt, err := base64.StdEncoding.DecodeString(a.Salt)
	if err != nil {
		return err
	}

	err = checkPassword([]byte(a.Password), []byte(old), salt)
	if err == nil {
		hash, err := hashPassword([]byte(newpass), salt)
		if err != nil {
			return err
		}
		a.Password = string(hash)
		a.UpdatedAt = time.Now()

		if err := s.DB.Update(&a); err != nil {
			return err
		}
		return nil

	} else {
		return baas.ErrBaasInvalidPassword
	}
}

// get private via email
func (s *BlockStorage) GetKey(ctx context.Context, id, password string) (string, string, string, string, error) {
	fmt.Printf("User try to get private key")
	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return "", "", "", "", baas.ErrBaasNoSuchUser
		}
		return "", "", "", "", err
	}
	salt, err := base64.StdEncoding.DecodeString(a.Salt)
	if err != nil {
		return "", "", "", "", err
	}
	//fmt.Println(a, password)
	err = checkPassword([]byte(a.Password), []byte(password), salt)
	if err == nil {
		return a.PrivateKey, a.CipherText, a.Rv, a.Salt, nil
	} else {
		return "", "", "", "", baas.ErrBaasInvalidPassword
	}

}
func (s *BlockStorage) DeleteKey(ctx context.Context, id, password string) error {
	fmt.Printf("User try to delete private key")
	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return baas.ErrBaasNoSuchUser
		}
		return err
	}
	salt, err := base64.StdEncoding.DecodeString(a.Salt)
	if err != nil {
		return err
	}
	//fmt.Println(a, password)
	err = checkPassword([]byte(a.Password), []byte(password), salt)
	if err == nil {
		err := s.DB.DeleteStruct(&a)
		return err
	} else {
		return baas.ErrBaasInvalidPassword
	}
}

func (s *BlockStorage) GetAddress(ctx context.Context, id, password string) (string, error) {
	fmt.Printf("User try to get address key")
	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return "", baas.ErrBaasNoSuchUser
		}
		return "", err
	}
	salt, err := base64.StdEncoding.DecodeString(a.Salt)
	if err != nil {
		return "", err
	}

	err = checkPassword([]byte(a.Password), []byte(password), salt)
	if err == nil {
		return a.Address, nil
	} else {
		return "", errors.New("Wrong password")
	}

}

// return application id based on userId/email
func (s *BlockStorage) GetApplicationId(ctx context.Context, id string) (string, error) {

	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return "", baas.ErrBaasNoSuchUser
		}
		return "", err
	}

	return a.ApplicationID, nil

}

// called by internal
func (s *BlockStorage) GetAddressByService(ctx context.Context, id, password string) (string, string, string, string, string, error) {

	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return "", "", "", "", "", baas.ErrBaasNoSuchUser
		}
		return "", "", "", "", "", err
	}

	if len(password) == 0 {
		return a.Address, "", "", "", "", nil
	}

	salt, err := base64.StdEncoding.DecodeString(a.Salt)
	if err != nil {
		return "", "", "", "", "", err
	}

	err = checkPassword([]byte(a.Password), []byte(password), salt)
	if err == nil {
		return a.Address, a.PrivateKey, a.Rv, a.CipherText, a.Salt, nil
	} else {
		return "", "", "", "", "", baas.ErrBaasInvalidPassword
	}

}

func (s BlockStorage) boltBaasUser(ctx context.Context, id string) (*blockUser, error) {
	var a blockUser
	if err := s.DB.One("Email", id, &a); err != nil {
		if err == storm.ErrNotFound {
			return nil, baas.ErrApplicationNotFound
		}
		return nil, err
	}

	return &a, nil
}

// user can update username and password, other like applicationid ,private key can't be changed
func (s *BlockStorage) Update(ctx context.Context, u *baas.UpdateBAASUser) (*baas.BAASUser, error) {
	fmt.Printf("User try to get private key")

	a, err := s.boltBaasUser(ctx, u.Email)
	if err != nil {
		return nil, err
	}

	a.UserName = u.UserName
	a.Password = u.Password // change username and password
	a.UpdatedAt = time.Now()

	if err := a.validate(s.DB); err != nil {
		return nil, err
	}

	if err := s.DB.Save(a); err != nil {
		if err == storm.ErrAlreadyExists {
			return nil, baas.ErrApplicationNameTaken
		}
		return nil, err
	}

	mahiApp := sanitizeblockUser(*a)

	return &mahiApp, nil

}
func sanitizeblockUser(t blockUser) baas.BAASUser {
	return baas.BAASUser{
		ID:            t.ID,
		Name:          t.UserName,
		ApplicationID: t.ApplicationID,
		Email:         t.Email,
		Address:       t.Address,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

// page start 1 表示第一页
func (s BlockStorage) paginateUsers(ctx context.Context, id string, page, limit int) ([]*baas.BAASUser, error) {
	var users []*blockUser
	var retUsers []*baas.BAASUser
	if limit == -1 {
		err := s.DB.Select(q.Eq("ApplicationID", id)).OrderBy("CreatedAt").Find(&users)
		if err != nil {
			if err == storm.ErrNotFound {
				return []*baas.BAASUser{}, nil
			}
			return nil, err
		}

	} else {
		realpage := 0
		if page >= 1 {
			realpage = page - 1
		}
		err := s.DB.Select(q.Eq("ApplicationID", id)).OrderBy("CreatedAt").Limit(limit).Skip(realpage * limit).Find(&users)
		if err != nil {
			if err == storm.ErrNotFound {
				return []*baas.BAASUser{}, nil
			}
			return nil, err
		}
	}
	for _, a := range users {
		aUser := sanitizeblockUser(*a)
		retUsers = append(retUsers, &aUser)
	}
	return retUsers, nil
}
func (s BlockStorage) ApplicationUsers(ctx context.Context, id string, page, limit int) ([]*baas.BAASUser, error) {
	return s.paginateUsers(ctx, id, page, limit)
}

// randSalt generates 16 * 8 bits of data for a random salt
func randSalt() ([]byte, error) {
	buf := make([]byte, 16)
	count := len(buf)
	n, err := crand.Read(buf)
	if err != nil {
		return nil, err
	}

	if n != count || err != nil {
		for count > 0 {
			count--
			buf[count] = byte(r.Int31n(256))
		}
	}

	return buf, nil
}

// saltPassword combines the salt and password provided
func saltPassword(password, salt []byte) ([]byte, error) {
	salted := &bytes.Buffer{}
	_, err := salted.Write(append(salt, password...))
	if err != nil {
		return nil, err
	}

	return salted.Bytes(), nil
}

// hashPassword encrypts the salted password using bcrypt
func hashPassword(password, salt []byte) ([]byte, error) {
	salted, err := saltPassword(password, salt)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword(salted, 10)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// checkPassword compares the hash with the salted password. A nil return means
// the password is correct, but an error could mean either the password is not
// correct, or the salt process failed - indicated in logs
func checkPassword(hash, password, salt []byte) error {
	salted, err := saltPassword(password, salt)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hash, salted)
}

package gm

import (
	"crypto/cipher"
	"encoding/hex"

	"github.com/agreyfox/baas/adapter/gm/util"
)

type EncryptionService struct {
	Key []byte
	Iv  []byte
}

func (s *EncryptionService) SetIV(iv []byte) {
	s.Iv = iv
}
func (s *EncryptionService) SetKey(key []byte) {
	s.Key = key
}

func (s *EncryptionService) Config(key []byte, iv []byte) {
	s.Key = key
	s.Iv = iv
}

/* 	use gm to encrypt text

encrypter := cipher.NewCBCEncrypter(c, data.iv)
		result := make([]byte, len(data.out))
		encrypter.CryptBlocks(result, util.PKCS5Padding(data.in, BlockSize))
		fmt.Printf("encrypt result:%s\n", hex.EncodeToString(result))
		if !bytes.Equal(result, data.out) {
			t.Error("encrypt result not equal expected")
			return
		}

*/
func (s *EncryptionService) Encrypt(plaintext []byte) ([]byte, error) {
	c, err := NewCipher(s.Key)
	if err != nil {
		return nil, err
	}

	gm := cipher.NewCBCEncrypter(c, s.Iv)

	m := util.PKCS5Padding(plaintext, BlockSize)
	result := make([]byte, len(m))
	gm.CryptBlocks(result, m)
	//fmt.Printf("encrypt result:%s\n", hex.EncodeToString(result))
	//fmt.Println("en:", result, len(result))
	return result, nil
}

func (s *EncryptionService) EncryptToString(plaintext []byte) (string, error) {
	ciphertext, err := s.Encrypt(plaintext)
	if err != nil {
		return "", err
	}

	//encoded := base64.URLEncoding.EncodeToString(ciphertext)
	encoded := hex.EncodeToString(ciphertext)
	return encoded, nil
}

/*


decrypter := cipher.NewCBCDecrypter(c, data.iv)
		plain := make([]byte, len(result))
		decrypter.CryptBlocks(plain, result)
		fmt.Printf("decrypt result:%s\n", hex.EncodeToString(plain))
		plain = util.PKCS5UnPadding(plain)
		if !bytes.Equal(plain, data.in) {
			t.Error("decrypt result not equal expected")
			return
		}
*/

func (s *EncryptionService) Decrypt(ciphertext []byte) ([]byte, error) {

	c, err := NewCipher(s.Key)
	if err != nil {
		return nil, err
	}
	//fmt.Println("enter", ciphertext, len(ciphertext))
	//fmt.Println(s.Key, s.Iv)

	gm := cipher.NewCBCDecrypter(c, s.Iv)
	plain := make([]byte, len(ciphertext))
	gm.CryptBlocks(plain, ciphertext)
	//fmt.Printf("decrypt:%s\n", hex.EncodeToString(plain))
	plain = util.PKCS5UnPadding(plain)

	return plain, nil
}

func (s *EncryptionService) DecryptString(ciphertext string) ([]byte, error) {
	decoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	return s.Decrypt(decoded)
}

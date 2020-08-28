package baas

import (
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
	"time"
)

var (
	r = mrand.New(mrand.NewSource(time.Now().Unix()))
)

func RandInt(min, max int) int {
	var seed int64
	binary.Read(crand.Reader, binary.BigEndian, &seed)

	mrand.Seed(seed)

	return mrand.Intn(max-min+1) + min
}

func RandStr(length int) string {
	var seed int64
	binary.Read(crand.Reader, binary.BigEndian, &seed)

	mrand.Seed(seed)
	alpha := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = alpha[mrand.Intn(len(alpha))]
	}
	return string(buf)
}

// randSalt generates 16 * 8 bits of data for a random salt
func GenerateSalt() ([]byte, error) {
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

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

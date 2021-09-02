package string_generator

import (
	b64 "encoding/base64"
	models2 "go-helper/models"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func EncryptBase64(stringToEncrypt string) (encryptedString string) {
	res := b64.StdEncoding.EncodeToString([]byte(stringToEncrypt))
	return res
}

func DecryptBase64(encryptedString string, keyString string) (decryptedString string,err error) {
	res, err := b64.StdEncoding.DecodeString(encryptedString)
	if err != nil{
		return "", models2.ErrGeneralMessage
	}
	return string(res),nil
}

func NowYmd(format string) string {
	t := time.Now()
	timeFormated := t.Format(format)
	return timeFormated
}



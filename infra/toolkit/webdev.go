package toolkit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func HmacSha256(s string) string {
	h := hmac.New(sha256.New, []byte("secret key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func EncodeBase64Custom(s string) string {
	return base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
}

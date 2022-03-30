package toolkit

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HmacSha256(t *testing.T) {
	t.Run("HmacSha256", func(t *testing.T) {
		expected := "c20a6b19e0ffa0f9d5c4004c032045be0b28a66946d90b438c7a967b4807995e"
		str := "aloysiusyoko@gmail.com"
		actual := HmacSha256(str)
		assert.Equal(t, expected, actual)
	})
}

func Test_EncodeBase64(t *testing.T) {
	t.Run("Base64 - Standard Encoding", func(t *testing.T) {
		expected := "eyJwaW5nIjogInBvbmcifQ=="
		str := `{"ping": "pong"}`
		actual := EncodeBase64(str)
		assert.Equal(t, expected, actual)

		// decode
		result, _ := base64.StdEncoding.DecodeString(actual)
		assert.Equal(t, str, string(result))
	})

	t.Run("Base64 - Custom Encoding", func(t *testing.T) {
		expected := "eyJwaW5nIjogInBvbmcifQ=="
		str := `{"ping": "pong"}`
		actual := EncodeBase64Custom(str)
		assert.Equal(t, expected, actual)
	})
}

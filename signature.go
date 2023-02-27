package go_jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func encodeSegment(seg []byte) string {
	return base64.RawURLEncoding.EncodeToString(seg)
}

func  sign(signingString string, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(signingString))

	return encodeSegment(hasher.Sum(nil))
}

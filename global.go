package go_jwt

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
)

var (
	ModfToken_ERR = errors.New("err :token已被篡改.The token has been tampered with")
	ExpirationTime_ERR = errors.New("err :token已过期.The token ExpirationTime")
)

var (
	Alg_HMAC_SHA1 = sha1.New
	Alg_HMAC_SHA256 = sha256.New
	Alg_HMAC_SHA384 = sha512.New384
	Alg_HMAC_SHA512 = sha512.New
	Alg_HMAC_RS256 = rsa.CRTValue{}
	Alg_HMAC_RS512= rsa.CRTValue{}


)
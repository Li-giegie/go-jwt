package go_jwt

import (
	"fmt"
	"testing"
)

type MyClaims struct {
	StandardClaims
}

type A struct {
	Sub  string `json:"sub,omitempty"`
	Name string `json:"name,omitempty"`
	Iat  int64  `json:"iat,omitempty"`
}

func (a A) Void()  {
}

func Test_Token(t *testing.T) {

	var token = Token{
		Header:    Header{
			Alg: "HS256",
			Typ: "JWT",
		},
		ClaimsI :&A{
			Sub: "1234567890",
			Name: "John Doe",
			Iat: 1516239022,
		},
	}
	var l = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.PjndlgESbR_7uPWq7tKFd6o7l799Y45mU5KvDcO2nPI"
	tokenStr ,err:= token.Marshal("12345678")
	fmt.Println(tokenStr == l,err,tokenStr)

	var token2 Token
	var a A
	fmt.Println(token2.Unmarshal(tokenStr,"12345678",&a),a)


}

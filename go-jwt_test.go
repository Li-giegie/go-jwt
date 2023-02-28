package go_jwt

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

type MyClaims struct {
	StandardClaims
	ASDF string
}



func Test_Token(t *testing.T) {

	var key = "黑猫警长"
	var token = Token{
		Header:    Header{
			Alg: "HS256",
			Typ: "JWT",
		},
		ClaimsI :&MyClaims{
			StandardClaims:StandardClaims{
				Iss: "",
				Sub: "",
				Aud: "",
				Exp: time.Now().UnixNano()+time.Second.Nanoseconds(),
				Nbf: 0,
				Iat: 0,
				Jti: "",
			},
			ASDF: "hello word !",
		},
	}
	//time.Sleep(time.Second)
	tokenStr,err := token.Marshal(key)
	fmt.Println("token : ",tokenStr,err)

	var myClaims MyClaims
	for  {
		err = Unmarshal(tokenStr,key,&myClaims)
		fmt.Println(err,myClaims)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}


}

func TestName(t *testing.T) {

	a,err := base64.RawURLEncoding.DecodeString(("eyJhIjoiMSJ9"))
	fmt.Println(string(a),err)
}
package go_jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Token struct {
	Header
	Payload
	ClaimsI
	signature
}
func (t *Token) Marshal(key string) (string,error) {
	err := t.Header.Marshal()
	err1 := t.Payload.Marshal(t.ClaimsI)
	if err != nil || err1 != nil {
		return "", err
	}

	t.signature.Marshal(t.Header.base64UrlCode,string(t.Payload),key)
	return string(t.signature),nil
}
func (t *Token) Unmarshal(tokenStr string,key string,obj interface{}) error {
	var err error
	_tmp := strings.Split(tokenStr,".")
	t.Header.base64UrlCode,t.Payload,t.signature = _tmp[0],Payload(_tmp[1]),signature(_tmp[2])
	if err = t.signature.Unmarshal(t.Header.base64UrlCode,string(t.Payload),key);err != nil {
		return err
	}

	if err = t.Header.Unmarshal(); err != nil {
		return err
	}

	return t.Payload.Unmarshal(obj)
}

type Header struct {
	Alg string `json:"alg,omitempty"`
	Typ string `json:"typ,omitempty"`
	base64UrlCode string
}

func (h *Header) Marshal() error  {
	hbuf,err := json.Marshal(h)
	if err != nil {
		return  err
	}
	h.base64UrlCode = base64.RawURLEncoding.EncodeToString(hbuf)
	return nil
}
func (h *Header) Unmarshal() error {
	buf,err := base64.RawURLEncoding.DecodeString(h.base64UrlCode)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf,h)
}

type Payload string

func (t *Payload) Marshal(c ClaimsI) error {
	buf,err := json.Marshal(c)
	if err != nil {
		return err
	}
	*t = Payload(base64.RawURLEncoding.EncodeToString(buf))
	return nil
}
func (t *Payload) Unmarshal(obj interface{}) error  {
	buf,err := base64.RawURLEncoding.DecodeString(string(*t))
	if err != nil {
		return  err
	}
	*t = Payload(buf)
	return json.Unmarshal(buf,obj)
}

type StandardClaims struct {
	Iss	string		// Issuer的简写，代表token的颁发者
	Sub string		// Subject的简写，代表token的主题
	Aud string		// Audience的简写，代表token的接收目标。
	Exp int64		// Expiration Time的简写，代表token的过期时间，时间戳格式。
	Nbf int64		// Not Before的简写，代表token在这个时间之前不能被处理，主要是纠正服务器时间偏差。
	Iat	int64		// Issued At的简写，代表token的颁发时间。
	Jti	string		// JWT ID的简写，代表token的id，通常当不同用户认证的时候，他们的token的jti是不同的。
}

func (s *StandardClaims) Void()  {
}
func (s *StandardClaims) GetExpirationTime() int64 {
	return s.Exp
}
type signature string
func (s *signature) Marshal(_Header,_Payload string,key string)  {
	*s = signature(_Header+"."+_Payload+"."+sign(_Header+"."+_Payload,key))
}
func (s *signature) Unmarshal(_Header,_Payload string,key string) error {
	if string(*s) != sign(_Header+"."+_Payload,key) {
		return ModfToken_ERR
	}
	return nil
}
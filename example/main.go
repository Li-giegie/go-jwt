package main

import (
	"fmt"
	jwt "github.com/Li-giegie/go-jwt"
	"log"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims
	Uid string
	Pwd string
}

// 颁发一个为期3秒钟的token,并解析它 直到过期为止
func main(){
	// 定义自己的需求这里是Uid、Pwd，可以内嵌jwt Claims标准字段，或者自己实现ClaimsI 接口，
	var myClaims = MyClaims{
		jwt.StandardClaims{
		Exp: time.Now().Add(time.Second*3).UnixNano(),
		Iat: time.Now().UnixNano(),
		},
		"admin", "123456",
	}
	// 创建token对象。header为固定值 目前仅支持hm256加密算法
	token := jwt.Token{
		Header:  jwt.Header{
			Alg: "hm256",
			Typ: "jwt",
		},
		ClaimsI: &myClaims,
	}
	// 定义一个Key用来加密
	var key = "i'm key ~"
	// 加密 入参是密钥key
	tokenStr,err := token.Marshal(key)
	if err != nil {
		log.Fatalln(err)
	}
	// 生成的token字符串
	log.Println("generate token :",tokenStr)


	// 定义一个待解析的对象 根据字符串生成 MyClaims需求
	var myClaims2 MyClaims
	// 解析token字符串，入参：token字符串、密钥、myClaims2
	err = jwt.Unmarshal(tokenStr,key,&myClaims2)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("analysis token :",myClaims2)

	// 不断解析我们的token字符串 直到它过期为止，过期时间我们在开始的时候定义为3秒
	var i int
	for {
		i++
		err = jwt.Unmarshal(tokenStr,key,&myClaims2)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(i," 解析成功：",myClaims2)
		time.Sleep(time.Second)
	}
}


package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2
//加盐
var mySecret = []byte("夏天夏天悄悄过去")

type MyClaims struct {
	Phone   int64  `json:"phone"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenToken(phone int64, name string) (string, error) {
	c := MyClaims{
		phone,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",                               //签发人
		},
	}
	fmt.Println(c)
	//使用指定的签名方法创建签名队形
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(token)
	return token.SignedString(mySecret)
	//使用指定的secret签名并获得完整编码后的token

}

func ParseToken(tokenString string)(*MyClaims,error)  {
	//解析token
	token,err := jwt.ParseWithClaims(tokenString,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret,nil
	})
	if err != nil{
		return nil, err
	}
	if claims,ok := token.Claims.(*MyClaims);ok && token.Valid{
		return claims, err
	}
	return nil, err
}
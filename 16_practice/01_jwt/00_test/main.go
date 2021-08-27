package main

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt"

	"time"
)

var jwtSecret=[]byte("test")

type  Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(username,password string)(string,error){
	nowTime :=time.Now()
	expireTime:=nowTime.Add(3*time.Hour)

	claims:=Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "gin-blog",
		},
	}
	//
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
	token,err:=tokenClaims.SignedString(jwtSecret)

	return token,err
}


// 验证token的函数
func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{},error){
		return jwtSecret,nil
	})

	if tokenClaims!=nil{
		if claims,ok:=tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid{
			return claims,nil
		}
	}
	//
	return nil,err
}

func main() {

	token,_:= GenerateToken("xingye-222","123456")
	fmt.Println("生成的token:",token)
	claim,err:=ParseToken(token)
	if err!=nil {
		fmt.Println("解析token出现错误：",err)
	}else if time.Now().Unix() > claim.ExpiresAt {
		fmt.Println("时间超时")
	}else {
		fmt.Println("username:",claim.Username)
		fmt.Println("password:",claim.Password)
	}

}
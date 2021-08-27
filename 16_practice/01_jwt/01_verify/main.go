package main

import (
	"crypto/rsa"
	"fmt"
	"log"
	"strings"
	"github.com/golang-jwt/jwt"
)

var (
	verifyKey *rsa.PublicKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Username string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

func main() {
	tokenString := strings.TrimSpace("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA0NzQ5NjAsIm9yaWdfaWF0IjoxNjI5ODcwMTYwLCJyb2xlIjoiY3VzdG9tZXIiLCJ1c2VyX2lkIjo5NywidXNlcm5hbWUiOiJiYW9saUB0dXJpbmd2aWRlby5uZXQiLCJ1dWlkIjoiYWE0NzRlMDEtMjY3OC00MDZhLThiMmEtYTY3OTFkZjQ5MGExIn0.do5FfhMfkeV3TX-INRz5CqpJ4CQmqBtpGzdjEhDV-aah2Edl28E677JXfQDQGvugrxrxie98M_r-PKwMFdEbYQ")

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	fatal(err)

	claims := token.Claims.(*CustomClaimsExample)
	fmt.Println(claims.CustomerInfo)

}


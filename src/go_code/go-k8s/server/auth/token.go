package auth

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type userStdClaims struct {
	StandardClaims jwt.StandardClaims `json:"standardClaims"`
	User           *User              `json:"user"`
}

type Token struct {
	TokenString string `json:"tokenString"`
}

var AppSecret = "abcdefghigklmn1234567890"
var gtoken = ""

func (claim userStdClaims) Valid() (err error) {
	if claim.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		log.Printf("token is expired.")
		return errors.New("token is expired")
	}
	if claim.StandardClaims.Id != "1" {
		return errors.New("invalid user Id")
	}
	return
}

func jwtGenToken(m *User, d time.Duration) (*Token, error) {
	m.Password = ""
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", 1),
		Issuer:    "kubernetes-cluster",
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		log.Printf("generate token failed")
	}
	gtoken = tokenString
	return &Token{TokenString: tokenString}, err
}

func parseToken(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claim := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	err = claim.Valid()
	if err != nil {
		return nil, err
	}
	return claim.User, err
}

func IsValidToken(htoken string) bool {
	// err := recover()
	if htoken == "" || !strings.HasPrefix(htoken, "Bearer ") {
		log.Panicln("header token is empty or is not start with Bearer")
		return false
	}
	htoken = strings.Split(htoken, " ")[1]

	htoken = strings.Trim(htoken, " ")
	// log.Println("request token:", htoken)
	// log.Println("register token:", gtoken)
	u, err := parseToken(htoken)
	if err != nil {
		log.Println("parse token failed")
		return false
	}
	if u.UserName != user.UserName {
		log.Printf("invalid user %s", u.UserName)
		return false
	}
	if gtoken == "" || gtoken != htoken {
		log.Printf("invalid token")
		return false
	}
	return true
}

// func main() {
// 	user := auth.User{
// 		UserName: "admin",
// 		Password: "123345",
// 	}
// 	result, err := jwtGenToken(&user, time.Second*5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	time.Sleep(time.Second * 6)
// 	u, err := parseToken(result)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(*u)
// 	}

// }

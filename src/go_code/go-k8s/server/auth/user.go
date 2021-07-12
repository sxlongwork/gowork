package auth

import (
	"errors"
	"fmt"
	"k8s-manage/server/utils"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var user = getUserInfo()

func GetToken(u *User) (*Token, error) {
	// log.Println(*u)
	// log.Println(user)
	// TODO
	if !u.verify() {
		return nil, errors.New("username or password is invalid\n")
	}
	log.Println("verify username and password success, generate token...")
	token, err := jwtGenToken(u, time.Hour*12)
	if err != nil {
		log.Println("generate token failed")
	} else {
		log.Println("generate token success")
	}
	return token, err
}

func getUserInfo() *User {
	umap := *utils.ConfMap
	username := umap["username"].(string)
	password := umap["password"].(string)
	if username == "" || password == "" {
		log.Fatal(fmt.Sprintf("init user %s failed", username))
	}
	log.Printf("init user %s success\n", username)
	return &User{UserName: username, Password: utils.Encode(password)}
}

func (u *User) verify() bool {
	// log.Println("request param", u.UserName, utils.Encode(u.Password))
	// log.Println("init user", user.UserName, user.Password)
	if u.UserName != user.UserName {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return false
	}
	return true
}

// func main() {
// 	fmt.Println(user)
// }

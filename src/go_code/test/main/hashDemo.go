package main

import (
	// "crypto/sha256"
	// "encoding/hex"
	// "log"
	"encoding/json"
	"fmt"
	// "github.com/gin-gonic/gin"
)

// func calculateHash(toBeHash string) string {
// 	hashInBytes := sha256.Sum256([]byte(toBeHash))
// 	hashInString := hex.EncodeToString(hashInBytes[:])
// 	log.Printf(format:"%s, %s", hashInBytes, hashInString)
// 	return hashInString
// }

type Cat struct {
	Name string `json:"name"`
}

func getName(cat interface{}) []byte {
	data, err := json.Marshal(cat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[]data =", data)
	// str := string(data)
	return data
}

func main() {
	// calculateHash("test01")

	//测试一下其他东西
	// fmt.Println(strconv.Itoa(2))

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, Geektutu")
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
	// var cat interface{}
	// data := getName(Cat{"tom"})
	// err := json.Unmarshal(data, cat)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cat)
	data := getName(Cat{"tom"})
	fmt.Println(string(data))
	newCat := Cat{}
	err := json.Unmarshal(data, &newCat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newCat)

	maptest := make(map[string]string)
	maptest["aa"] = "001"
	maptest["bb"] = "002"
	for k, v := range maptest {
		fmt.Println(k, v)
	}

}

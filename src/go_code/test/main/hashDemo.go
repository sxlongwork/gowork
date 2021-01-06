package main

import (
	// "crypto/sha256"
	// "encoding/hex"
	// "log"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)


// func calculateHash(toBeHash string) string {
// 	hashInBytes := sha256.Sum256([]byte(toBeHash))
// 	hashInString := hex.EncodeToString(hashInBytes[:])
// 	log.Printf(format:"%s, %s", hashInBytes, hashInString)
// 	return hashInString
// }

func main() {
	// calculateHash("test01")

	//测试一下其他东西
	fmt.Println(strconv.Itoa(2))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	
}
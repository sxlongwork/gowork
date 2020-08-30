package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)


func calculateHash(toBeHash string) string {
	hashInBytes := sha256.Sum256([]byte(toBeHash))
	hashInString := hex.EncodeToString(hashInBytes[:])
	log.Printf(format:"%s, %s", hashInBytes, hashInString)
	return hashInString
}

func main() {
	calculateHash("test01")
}
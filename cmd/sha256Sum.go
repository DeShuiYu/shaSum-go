package main

import (
	sha2562 "crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func sha256(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	hash := sha2562.New()
	io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Failed and need to provide a file path")
	}
	for _, filePath := range args[1:] {
		fmt.Printf("%v\t%s\n", filePath, sha256(filePath))
	}
}

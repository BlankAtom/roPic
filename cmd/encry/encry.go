package main

import (
	"fmt"
	"github.com/blackswords/roPic/internal/entrypt"
	"os"
)

func main() {

	//s := "123456"
	//t := entrypt.Encrypt(s)
	//fmt.Printf(t)
	//t = entrypt.Decrypt(t)
	//fmt.Printf(t)
	argsLen := len(os.Args)
	if argsLen < 2 {
		fmt.Printf("Didnt has more parameters.")
		return
	}

	isDecrypt := os.Args[1]
	var res string
	if isDecrypt == "de" || isDecrypt == "decrypt" {
		if argsLen < 3 {
			fmt.Printf("Need third param.")
			return
		}
		res = entrypt.Decrypt(os.Args[2])
	} else {
		res = entrypt.Encrypt(isDecrypt)
	}

	fmt.Printf(res)

	//for i := 0; i < args_len; i++ {
	//	wait_encry_string := os.Args[i]
	//	// Encry
	//}
	//secret_key := flag.String("sk", "", "Secret key.")
	//secret_id := flag.String("si", "", "Secret ID.")
	//app_id := flag.String("a", "", "App ID.")
	//region := flag.String("r", "", "region.")
	//bucket_name := flag.String("bn", "", "Bucket name.")
	//version := flag.Int("v", 5, "bucket version.")
	//path := flag.String("p", "", "Save path in bucket.")

}

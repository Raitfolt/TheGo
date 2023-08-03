package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	str1 = "string1"
	str2 = "string2"
)

func main() {

	method := flag.String("method", "SHA256", "method of hash calculation(SHA256, SHA384, SHA512)")
	flag.Parse()

	var result1, result2 string

	switch *method {
	case "SHA256":
		fmt.Println("SHA256")
		sha1 := sha256.Sum256([]byte(str1))
		sha2 := sha256.Sum256([]byte(str2))
		result1 = fmt.Sprintf("%x\n", sha1)
		result2 = fmt.Sprintf("%x\n", sha2)
	case "SHA384":
		fmt.Println("SHA384")
		sha1 := sha512.Sum384([]byte(str1))
		sha2 := sha512.Sum384([]byte(str2))
		result1 = fmt.Sprintf("%x\n", sha1)
		result2 = fmt.Sprintf("%x\n", sha2)
	case "SHA512":
		fmt.Println("SHA512")
		sha1 := sha512.Sum512([]byte(str1))
		sha2 := sha512.Sum512([]byte(str2))
		result1 = fmt.Sprintf("%x\n", sha1)
		result2 = fmt.Sprintf("%x\n", sha2)
	}

	fmt.Println(result1)
	fmt.Println(result2)
}

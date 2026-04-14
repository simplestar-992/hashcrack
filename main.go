package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"os"
)

var (
	algorithm = flag.String("a", "sha256", "Hash algorithm: md5, sha1, sha256, sha512")
	compare   = flag.String("c", "", "Compare against hash")
	fileFlag  = flag.String("f", "", "Hash a file")
)

func main() {
	flag.Parse()

	var h hash.Hash

	switch *algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		fmt.Printf("❌ Unknown algorithm: %s\n", *algorithm)
		os.Exit(1)
	}

	var data []byte

	if *fileFlag != "" {
		f, err := os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		data = f
	} else {
		// Read from stdin or args
		if flag.NArg() > 0 {
			data = []byte(flag.Arg(0))
		} else {
			fmt.Scan(&data)
		}
	}

	h.Write(data)
	result := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("%s: %s\n", *algorithm, result)

	if *compare != "" {
		if result == *compare {
			fmt.Println("✅ MATCH!")
		} else {
			fmt.Println("❌ NO MATCH")
		}
	}
}


package main

import (
	"hash/crc32"
	"fmt"
	"os"
	"io"
	"crypto/sha1"
)

func main()  {
	//non-crypto hashes

	// common use of crc32 is to compare two files
	hasher := crc32.NewIEEE()

	hasher.Write([]byte("test"))

	value := hasher.Sum32()

	fmt.Println(value)

	hashOne, err := getHash("io.txt")
	if err != nil {
		return
	}
	hashTwo, err := getHash("testio.txt")
	if err != nil {
		return
	}
	fmt.Println("Result : ", hashOne, hashTwo, hashOne == hashTwo)


	// crypto hashes
	cryptoHasher := sha1.New()
	cryptoHasher.Write([]byte("test"))
	bs := cryptoHasher.Sum([]byte{})
	fmt.Println("encrypted: ", bs)
	fmt.Println("hasher: ", cryptoHasher)

}

func getHash(filename string) (uint32, error)  {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	hasher := crc32.NewIEEE()

	_, err = io.Copy(hasher, file)

	if err != nil {
		return 0, err
	}
	return hasher.Sum32(), nil

}
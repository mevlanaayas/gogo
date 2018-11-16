
package main

import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"errors"
)

func main()  {
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf.Bytes())

	file, err := os.Open("io.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// defer closes file when main finished
	// so i did it manually to continue now
	file.Close()

	// following line is the common way to read file
	bs, err = ioutil.ReadFile("io.txt")
	if err != nil {
		return
	}

	str = string(bs)
	fmt.Println(str)
	file.Close()

	// file creation
	file, err = os.Create("creation.txt")
	if err != nil {
		return
	}
	file.WriteString("test file creation")

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo)
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	err = errors.New("error message")
	fmt.Println(err)

}
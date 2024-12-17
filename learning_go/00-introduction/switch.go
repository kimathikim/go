package main

import (
	"fmt"
	"time"
	"os"
)

type File struct {
	*os.File
}

func (file *File) Write(b []byte) (n int, err error) {
	return file.File.Write(b)
}

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

myFile := &File{file}
	_, err = myFile.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	today := time.Now()

	fmt.Println(today.Day())

	switch today.Day() {
	case 5:
		fmt.Println("wash your clothes")
		fallthrough
	case 13, 17:
		fmt.Println("just relax buddy!!!")
		fallthrough
	default:
		fmt.Println("never mind. It is not worth it")
	}
}

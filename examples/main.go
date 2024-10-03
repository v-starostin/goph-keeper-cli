package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("start")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("wd:", wd)
	file, err := os.Open("./examples/img.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	resultFile, err := os.Create("examples/resultfile.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resultFile.Close()

	buf := make([]byte, 200)

	var bytesHaveRead int
	var bytesHaveWritten int
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Printf("have read %d bytes\n", n)
		bytesHaveRead += n
		n, err = resultFile.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Printf("have written %d bytes\n", n)
		bytesHaveWritten += n
	}
	fmt.Println("finish")
	fmt.Printf("%d bytes have been read\n", bytesHaveRead)
	fmt.Printf("%d bytes have been written\n", bytesHaveWritten)
}

package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, err := os.Open("./readme")
	if err != nil {
		fmt.Println("open file failed !, err :", err)
	}

	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了！")
			break
		}
		if err != nil {
			fmt.Println("read file failed , err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

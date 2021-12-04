package iniParse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

//ini配置文件解析器

//MysqlConfig 配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(config *MysqlConfig) {
	file, err := os.Open("./conf.ini")
	if err != nil {
		fmt.Println("open file failed !, err :", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				index := strings.Index(line, "=")
				if index > -1 {
					key := line[:index]
					value := line[index+1:]
					fmt.Printf("key:%s,value:%s\n", key, value)
					switch key {
					case "address":
						config.Address = value
					case "port":
						i, _ := strconv.Atoi(value)
						config.Port = i
					case "username":
						config.Username = value
					case "password":
						config.Password = value
					}
				}
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}

		if len(line) != 0 {
			index := strings.Index(line, "=")
			if index > -1 {
				key := strings.ToLower(line[:index])
				value := strings.Replace(line[index+1:], "\n", "", -1)
				value = strings.Replace(value, "\r", "", -1)
				fmt.Printf("key:%s,value:%s\n", key, value)
				switch key {
				case "address":
					config.Address = value
				case "port":
					i, _ := strconv.Atoi(value)
					config.Port = i
				case "username":
					config.Username = value
				case "password":
					config.Password = value
				}
			}
		}
	}
}

func TestParse(t *testing.T) {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Printf("MysqlConfig: %s, %d, %s, %s\n", mc.Address, mc.Port, mc.Username, mc.Password)
}

package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "http://127.0.0.1:9090/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed, err:", err)
		return
	}
	fmt.Println(string(bytes))

}

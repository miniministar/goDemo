package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "hello, GoLang!")
}

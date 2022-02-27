package main

import (
	"httpserver/utils"
	"net/http"
)

// http协议 http://localhost:9090/testpath/1/

func main() {

	http.HandleFunc("/testpath/1/", utils.HandlerFunc)
	http.ListenAndServe("127.0.0.1:9090", nil)

}

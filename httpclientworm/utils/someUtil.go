package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHttp() {
	resp, err := http.Get("http://127.0.0.1:9090/testpath/1/")
	if err != nil {
		fmt.Printf("error : ", err)
		return
	}
	fmt.Println("resp.Request = ", resp.Request)
	fmt.Println("resp.Header = ", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error : ", err)
		return
	}
	fmt.Println("resp.Body = ", b)

}

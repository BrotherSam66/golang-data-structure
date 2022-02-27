package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandlerFunc 回应
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Method = ", r.Method)
	bBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("r.Body = %v \n", bBody)
	fmt.Println("r.Body = ", r.Body)
	str := `<h1>测试一级标题</h1><h3>测试3级标题</h3>`
	w.Write([]byte(str))
	b, err := ioutil.ReadFile("./page.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandlerFunc 回应
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.URL = ", r.URL)
	queryParam := r.URL.Query() // 自动识别URL里面的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println("name = ", name, " ;  age = ", age)
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.Header = ", r.Header)
	bBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("r.Body = %v \n", bBody)
	fmt.Println("r.Body = ", string(bBody))
	fmt.Println("r.Body = ", r.Body)
	str := `<h1>测试一级标题</h1><h3>测试3级标题</h3>`
	w.Write([]byte(str))
	b, err := ioutil.ReadFile("./page.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

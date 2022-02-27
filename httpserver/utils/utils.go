package utils

import "net/http"

// HandlerFunc 回应
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	str := `<h1>测试一级标题</h1><h3>测试3级标题</h3>`
	w.Write([]byte(str))
}

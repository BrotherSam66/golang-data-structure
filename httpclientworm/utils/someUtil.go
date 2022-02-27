package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HttpGet 两种get请求
func HttpGet() {
	//resp, err := http.Get("http://127.0.0.1:9090/testpath/1/")
	resp, err := http.Get(`http://127.0.0.1:9090/testpath/1/?name=MYNAME&age=18`)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer resp.Body.Close() // 最后要关闭连接。放在err处理后面就对了
	fmt.Println("resp.Request = ", resp.Request)
	fmt.Println("resp.Header = ", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	//fmt.Println("resp.Body = ", resp.Body)
	fmt.Println("resp.Body = ", string(b))

	// http.NewRequest() 方法，可以拼接url
	data := url.Values{} // url encode url编码
	data.Set("name", "周琳?&name=周琳")
	data.Set("age", "18")
	queryStr := data.Encode() // url编码后的值
	fmt.Println("queryStr =: ", queryStr)

	urlObj, _ := url.Parse("http://127.0.0.1:9090/testpath/1/")
	urlObj.RawQuery = queryStr                                  // 这个干什么的？
	reqObj, err := http.NewRequest("GET", urlObj.String(), nil) // nil是body的位置
	resp_2, err := http.DefaultClient.Do(reqObj)                // 实际上是建立一个连接
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer resp_2.Body.Close() // 最后要关闭连接。放在err处理后面就对了
	fmt.Println("resp.Request = ", resp_2.Request)
	fmt.Println("resp.Header = ", resp_2.Header)
	b, err = ioutil.ReadAll(resp_2.Body)
	if err != nil {
		fmt.Printf("error : ", err)
		return
	}
}

// HttpGetAddHeader 向HEADER加参数，仿真各种浏览器环境
func HttpGetAddHeader() {
	//client := &http.Client{CheckRedirect: redirectPolicyFunc}
	// 设定短连接（禁止长连接）。否则，5秒拉取一次，连接就半个小时满了。（）
	// 另外的办法是client用同一个，把client放到全局变量里面。（适用于连接频繁，几秒一次的那种，会不会被被拉服务器发现了啊）
	tr := &http.Transport{DisableKeepAlives: true}
	// 创建客户端
	client := &http.Client{Transport: tr}
	resp, err := http.Get(`http://127.0.0.1:9090/testpath/1/`)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer resp.Body.Close() // 最后要关闭连接。放在err处理后面就对了
	fmt.Println("resp.Request = ", resp.Request)
	fmt.Println("resp.Header = ", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	fmt.Println("resp.Body = ", string(b))

	// 加HEADER参数
	reqObj, err := http.NewRequest("GET", "http://127.0.0.1:9090/testpath/1/", nil) // nil是body的位置
	reqObj.Header.Add("If-None-Match", `W/"wyzzy"`)
	reqObj.Header.Add("xxxxx", `yyyyyyyy`)

	resp_2, err := client.Do(reqObj) // 实际上是建立一个连接
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer resp_2.Body.Close() // 最后要关闭连接。放在err处理后面就对了
	fmt.Println("resp.Request = ", resp_2.Request)
	fmt.Println("resp.Header = ", resp_2.Header)
	b, err = ioutil.ReadAll(resp_2.Body)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
}

// HttpPost 上传文件请求
func HttpPost() {
	//resp, err := http.Post("http://127.0.0.1:9090/testpath/post/", "image/jpeg", &buf)
	//if err != nil {
	//	fmt.Println("error : ", err)
	//	return
	//}
	//defer resp.Body.Close() // 最后要关闭连接。放在err处理后面就对了

}

// HttpPostForm 上传表单请求
func HttpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:9090/testpath/1/", url.Values{"mykey": {"myvalue"}, "age": {"18"}})
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer resp.Body.Close() // 最后要关闭连接。放在err处理后面就对了
	fmt.Println("resp.Request = ", resp.Request)
	fmt.Println("resp.Header = ", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	//fmt.Println("resp.Body = ", resp.Body)
	fmt.Println("resp.Body = ", string(b))

}

package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func CloneHeader(src http.Header, desc *http.Header) {
	for k, v := range src {
		desc.Set(k, v[0])
	}
}

func RequestUrl(w http.ResponseWriter, r *http.Request, url string) {

	r.Header.Set("x-forwarded-for", r.RemoteAddr)
	newReq, _ := http.NewRequest(r.Method, url, r.Body)
	CloneHeader(r.Header, &newReq.Header)
	resp, _ := http.DefaultClient.Do(newReq)

	defer resp.Body.Close()
	//获取当前头信息
	getHeader := w.Header()
	//将请求回来的头信息赋值给当前头信息
	CloneHeader(resp.Header, &getHeader)
	//将请求回来的状态码给当前的头
	w.WriteHeader(resp.StatusCode)

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("请求来自：%s\n", r.RemoteAddr)
	w.Write(content)
}

func Catch() {
	if err := recover(); err != nil {
		for i := 3; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}

			log.Println(pc, file, line)
		}
	}
}

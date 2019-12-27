package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type Web1 struct {
}

func (web Web1) GetIp(req *http.Request) string {
	ips := req.Header.Get("x-forwarded-for")
	if ips != "" {
		ips_list := strings.Split(ips, ",")
		if len(ips_list) >0 && ips_list[0] != "" {
			return ips_list[0]
		}
	}

	return req.RemoteAddr
}

func (web Web1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Secure Area\"")
		w.WriteHeader(401)
		return
	}

	auths := strings.Split(auth, " ")
	if (len(auths) == 2) && (auths[0] == "Basic") {
		b,_ := base64.StdEncoding.DecodeString(auths[1])
		if string(b) == "admin:123" {
			fmt.Fprintf(w, "请求来自：%s", web.GetIp(r))
			return
		}
	}

	w.Write([]byte("用户名密码错误"))
}

type Web2 struct {
}

func (web Web2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("web2"))
}

type Web3 struct {
}

func (web Web3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("web3"))
}

func main()  {

	c := make(chan os.Signal)

	go func() {
		log.Fatal(http.ListenAndServe(":9091", Web1{}))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":9092", Web2{}))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":9093", Web3{}))
	}()

	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println(s)
}

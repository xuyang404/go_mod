package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"network/utils"
	"regexp"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer utils.Catch()

	for k, v := range utils.ProxyConfigs {
		if matched, _ := regexp.MatchString(k, r.URL.Path); matched == true {
			target,err := url.Parse(v)
			if err != nil{
				log.Println(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(w,r)
			//utils.RequestUrl(w, r, v)
			return
		}
	}

	w.Write([]byte("default"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &Proxy{}))
}

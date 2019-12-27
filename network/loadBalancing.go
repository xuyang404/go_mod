package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"network/utils"
)

type Proxy2 struct {

}

func (this Proxy2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer utils.Catch()

	if r.URL.Path == "/favicon.ico" {
		return
	}
	//server,err := url.Parse(ld.SelectServerByRand().Host)
	server,err := url.Parse(utils.LB.SelectServerByWeightPolling3().Host)
	if err != nil {
		log.Panicln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(server)
	proxy.ServeHTTP(w, r)
	return
}

func main()  {
	log.Fatal(http.ListenAndServe(":8081", &Proxy2{}))
}

package main

import (
	"GinHello/pkg/setting"
	"GinHello/routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	router := routers.InitRouter()

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	server := endless.NewServer(fmt.Sprintf(":%d", setting.HTTPPort), router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	log.Fatal(server.ListenAndServe())
}

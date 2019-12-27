package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main()  {
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	port := conf.GetIntDefault("app.service.port", 18001)
	fmt.Println(port)
}

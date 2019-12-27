package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var ProxyConfigs map[string]string
var Cfg *ini.File

func init()  {
	ProxyConfigs = make(map[string]string)
	var err error
	Cfg, err = ini.Load("env")
	if err != nil {
		fmt.Println(err)
	}

	sec, err := Cfg.GetSection("proxy")
	if err != nil {
		fmt.Println(err)
	}

	if sec != nil {
		secs := sec.ChildSections()
		for _,sec := range secs{

			pass,_ := sec.GetKey("pass")
			path,_ := sec.GetKey("path")
			if pass != nil && path != nil {
				ProxyConfigs[pass.Value()] = path.Value()
			}
		}
	}
}



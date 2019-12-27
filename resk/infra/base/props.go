package base

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(cxt infra.StarterContext)  {
	props = ini.NewIniFileConfigSource("config.ini")
}
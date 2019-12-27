package internal

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

var (
	Json jsoniter.API
)

type Route struct {
	Url  string    `json:"url"`
	Time time.Time `json:"time"`
}

func init() {
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
}

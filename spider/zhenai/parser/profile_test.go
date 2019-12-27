package parser

import (
	"fmt"
	"testing"
)



//用户信息解析器
func TestPro(t *testing.T) {
	strs := []string{"abc", "def"}
	count := len(strs)
	a := 5 - count
	for i:=0; i<=a;i++  {
		strs = append(strs, "")
	}
	fmt.Println(len(strs))
}


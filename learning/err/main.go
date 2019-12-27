package main

import (
	"fmt"
	"runtime"
)

func stupidCode() {
	n := 0
	fmt.Println(1 / n)
}

func catch() {
	if err := recover(); err != nil {
		for i := 3; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Println(pc, file, line)
		}
	}
}

func main() {
	defer catch()

	stupidCode()
}

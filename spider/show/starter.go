package main

import (
	"net/http"
	"spider/show/controller"
)

func main()  {
	http.Handle("/search",
		controller.CreateSearchResultHandle("./show/view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

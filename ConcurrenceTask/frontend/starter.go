package main

import (
	"GoSpider/ConcurrenceTask/frontend/controller"
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/search", controller.SearchResultHandler{})
	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("listen:http://localhost:8888")
}

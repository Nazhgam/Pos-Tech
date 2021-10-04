package main

import (
	"log"
	"net/http"
	"pos-tech/helper"
)

func main() {
	http.HandleFunc("/get/id", helper.GetPostById) // post/ id:int
	http.HandleFunc("/create", helper.CreatePost)  // post/ name, task, author
	http.HandleFunc("/update/id", helper.Update)   //
	http.HandleFunc("/delete/id", helper.Delete)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("server have error", err)
		return
	}
}

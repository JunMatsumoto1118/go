package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	
	var hoge string = r.FormValue("msg")

	fmt.Fprint(w, "<font size ="3">Hello, HTTPサーバ")
	fmt.Fprint(w, "\n", hoge)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

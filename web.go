package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var hoge string = r.FormValue("msg")

	fmt.Fprint(w, "Hello, HTTPサーバ")
	fmt.Fprint(w, "\n", hoge)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

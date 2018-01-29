package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func unsei(num int) string {
	var message string = ""
	switch num {
	case 1:
		message = "凶"
	case 2, 3:
		message = "吉"
	case 4, 5:
		message = "中吉"
	case 6:
		message = "大吉"
	}
	return message
}

func handler(w http.ResponseWriter, r *http.Request) {

	var omikuji int = 0
	var daikichi int = 0
	rand.Seed(time.Now().UnixNano())

	for daikichi <= 99 {
		var num int = rand.Intn(6) + 1
		var result string = unsei(num)
		if num == 6 {
			daikichi = daikichi + 1
		}
		omikuji = omikuji + 1
		fmt.Fprint(w, omikuji, "回目: ", result, "\n")
	}

	fmt.Fprint(w, "おみくじは", omikuji, "回回しました\n")
	fmt.Fprint(w, "大吉の出現回数は", daikichi, "回です\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

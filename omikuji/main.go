package main

import (
	"fmt"
	"math/rand"
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

func main() {
	var daikichi int = 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 100; i = i + 1 {
		var num int = rand.Intn(6) + 1
		var result string = unsei(num)

		if num == 6 {
			daikichi = daikichi + 1
		}
		fmt.Printf("%v : %v\n", num, result)
	}
	fmt.Printf("大吉の出現回数は%v回です", daikichi)
}

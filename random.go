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
		fmt.Printf("%v 回目: %v\n", omikuji, result)
	}

	fmt.Printf("おみくじの実施は%v回回しました\n", omikuji)
	fmt.Printf("大吉の出現回数は%v回、出現確率は%v%%です\n", daikichi, float32(daikichi)/float32(omikuji)*100)
}

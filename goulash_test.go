package main

import(
	"fmt"
	"goulash"
)

func main() {
	bot := goulash.BaseBot
	bot.ResponseFunc = func(text string) {
		fmt.Println(text)
	}
	bot.Run()
}

package goulash

import(
	"testing"
	"fmt"
)

func TestBotInit(t *testing.T) {
	bot := goulash.BaseBot
	bot.ResponseFunc = func(text string) {
		fmt.Println(text)
	}
	if bot == nil {
		t.Error("Failure to construct.")
	}
}

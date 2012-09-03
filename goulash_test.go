package goulash

import(
	"testing"
	"fmt"
)

func TestBotInit(t *testing.T) {
	responseFunc := func(text string) {
		fmt.Println(text)
	}
	bot := New("irc.freenode.net", "#go-nuts", "goulashBot", responseFunc)
	if bot == nil {
		t.Error("Failure to construct.")
	}
}

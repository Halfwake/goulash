package goulash

import(
	"testing"
	"fmt"
)

func TestBotInit(t *testing.T) {
	responseFunc := func(text string) {
		fmt.Println(text)
	}
	startFunc := func() {
	}
	bot := New("irc.freenode.net", "#go-nuts", "goulashBot", responseFunc, startFunc)
	if bot == nil {
		t.Error("Failure to construct.")
	}
}

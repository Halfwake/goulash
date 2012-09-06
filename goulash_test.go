package goulash

import(
	"testing"
	"log"
)

func TestBotInit(t *testing.T) {
	bot := New("irc.freenode.net", "#go-nuts", "goulashBot")
	if bot == nil {
		log.Fatal("Failure to construct.")
	}
	bot.OnStart = func(bot *BaseBot) {}
	bot.Response = func(bot *BaseBot, text string) {bot.Quit()}
	bot.Run()
}

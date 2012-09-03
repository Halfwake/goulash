package goulash

import (
	"fmt"
	"net"
)

type BaseBot struct {
	server, channel, botnick string
	conn net.Conn
	MainLoop func()
}

func NewBaseBot(server, channel, botnick string, main_loop_func func()) *BaseBot {
	var new_obj BaseBot
	new_obj.server = server
	new_obj.channel = channel
	new_obj.botnick = botnick
	new_obj.Connect()
	new_obj.MainLoop = main_loop_func	
	return &new_obj
}

func (bot *BaseBot) Connect() {
	temp_con, err := net.Dial("tcp",bot.server + ":" + "6667")

	if err != nil {
		//handle error
	} else {
		bot.conn = temp_con
	}
}

func (bot *BaseBot) Send(msg string){
	fmt.Fprintf(bot.conn, msg)
}

func (bot *BaseBot) Ping(target string) {
	bot.Send("PING" + " " + target)
}

func (bot *BaseBot) SendMsg(msg string) {
	bot.Send("PRIVMSG " + bot.channel + " :" + msg + "\n")
}

func (bot *BaseBot) JoinChannel() {
	bot.Send("JOIN" + " " + bot.channel + "\n")
}

func (bot *BaseBot) Quit() {
	bot.Send("QUIT\n")
}

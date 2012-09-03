package goulash

import (
	"fmt"
	"net"
)

type BaseBot struct {
	server, channel, botnick string
	conn net.Conn
	ResponseFunc func(string)
}

func NewBaseBot(server, channel, botnick string, response_func func(string)) *BaseBot {
	var new_obj BaseBot
	new_obj.server = server
	new_obj.channel = channel
	new_obj.botnick = botnick
	new_obj.Connect()
	new_obj.ResponseFunc = response_func	
	return &new_obj
}

func (bot *BaseBot) Run() {
	var text, requester string
	var ping_flag bool
	var buffer string
	for {
		text = bot.Recv(&buffer)
		ping_flag = bot.GetPing(text)
		if ping_flag {
			requester = bot.GetPingName(buffer)
			bot.Pong(requester)
		} else {
			bot.ResponseFunc(text)
		}
	}
}

func (bot *BaseBot) GetPingName(text string) string {
	return text[5:]
}

func (bot *BaseBot) GetPing(text string) bool {
	if text[:4] == "PING" {
		return true
	} else {
		return false
	}
}

func (bot *BaseBot) Connect() {
	temp_con, err := net.Dial("tcp", bot.server + ":" + "6667")

	if err != nil {
		//handle error
	} else {
		bot.conn = temp_con
	}
}

func (bot *BaseBot) Recv(buffer string) {
	_, err := bot.conn.Read(string)	
	if err != nil {
		//handle error
	}
}
func (bot *BaseBot) Send(msg string) {
	fmt.Fprintf(bot.conn, msg)
}

func (bot *BaseBot) Ping(target string) {
	bot.Send("PING" + " " + target)
}

func (bot *BaseBot) Pong(target string) {
	bot.Send("PONG" + " " + target)
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

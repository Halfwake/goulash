package goulash

import (
	"fmt"
	"net"
)

type baseBot struct {
	server, channel, botnick string
	conn net.Conn
	ResponseFunc func(string)
}

func New(server, channel, botnick string, responseFunc func(string)) *baseBot {
	var new_obj baseBot
	new_obj.server = server
	new_obj.channel = channel
	new_obj.botnick = botnick
	new_obj.connect()
	new_obj.ResponseFunc = responseFunc	
	return &new_obj
}

//Starts the main loop of the bot. Only pings are handled by default.
func (bot *baseBot) Run() {
	var text, requester string
	var pingFlag bool
	for {
		text = bot.Recv()
		pingFlag = bot.isPing(text)
		if pingFlag {
			requester = bot.pingName(text)
			bot.Pong(requester)
		} else {
			bot.ResponseFunc(text)
		}
	}
}

//Gets the name of who sent a ping request.
func (bot *baseBot) pingName(text string) string {
	return text[5:]
}

//Decides whether or not a string is a ping request.
func (bot *baseBot) isPing(text string) bool {
	return text[:4] == "PING"
}

//Connects the bot to whatever server is found in its 'server' field.
func (bot *baseBot) connect() {
	tempCon, err := net.Dial("tcp", bot.server + ":" + "6667")

	if err != nil {
		//handle error
	} else {
		bot.conn = tempCon
	}
}

//Connect the bot to a target server.
func (bot *baseBot) Connect(server string) {
	bot.server = server
	bot.connect()
}

//Recovers text from the 'conn' field of the bot.
func (bot *baseBot) Recv() string {
	var buffer []byte
	_, err := bot.conn.Read(buffer)	

	if err != nil {
		//handle error
	}
	return string(buffer)
}

//Sends a string through the 'conn' field of the bot.
func (bot *baseBot) Send(msg string) {
	fmt.Fprintf(bot.conn, msg)
}

//Pings a target.
func (bot *baseBot) Ping(target string) {
	bot.Send("PING" + " " + target)
}

//Pongs a target, is handled automatically by the 'Run' method.
func (bot *baseBot) Pong(target string) {
	bot.Send("PONG" + " " + target)
}

//Sends a private message to a target.
func (bot *baseBot) SendMsg(msg string) {
	bot.Send("PRIVMSG " + bot.channel + " :" + msg + "\n")
}

//Joins the channel found in the bots 'channel' field.
func (bot *baseBot) JoinChannel() {
	bot.Send("JOIN" + " " + bot.channel + "\n")
}

//Quits the server.
func (bot *baseBot) Quit() {
	bot.Send("QUIT\n")
}

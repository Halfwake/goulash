package goulash

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	PrivMatch = regexp.MustCompile(`^(:\w+\s+)?PRIVMSG\s+\w\s+`)
	PingMatch = regexp.MustCompile(`^PING\s+\w+\s*`)
	PongMatch = regexp.MustCompile(`^PONG\s+\w+\s*`)
	//Needs some server startup regexp matches
)

type baseBot struct {
	server, channel, botNick string
	conn net.Conn
	ResponseFunc func(string)
	StartFunc func()
}

func New(server, channel, botNick string, responseFunc func(string), startFunc func()) *baseBot {
	var newObj baseBot
	newObj.server = server
	newObj.channel = channel
	newObj.botNick = botNick
	newObj.ResponseFunc = responseFunc	
	newObj.StartFunc = startFunc
	return &newObj
}

//Starts the main loop of the bot. Only pings are handled by default.
func (bot *baseBot) Run() {
	var text, requester string

	bot.StartFunc()
	if bot.conn == nil {bot.Connect("")} //If the 'StartFunc' doesn't connect the bot it will connect automatically
	//The text processing loop
	for {
		text = bot.Recv()
		if bot.isPing(text) { //Handles all ping messages.
			requester = bot.pingName(text)
			bot.Pong(requester)
		} else if text != "" { //Handles all text. If the text does not contain any characters it is ignored.
			bot.ResponseFunc(text)
		}
	}
}

//Gets the name of who sent a ping request.
func (bot *baseBot) pingName(text string) string {
	return strings.TrimSpace(text[5:])
}

//Decides whether or not a string is a ping request.
func (bot *baseBot) isPing(text string) bool {
	return ""  != PingMatch.FindString(text)
}

//Connects the bot to the target server, and changes the server field  If the target server is an empty string the bot
//will attempt to connect to its 'server' field.
func (bot *baseBot) Connect(server string) {
	if server != "" {bot.server = server}
	tempCon, err := net.Dial("tcp", bot.server + ":" + "6667")

	if err != nil {
		//handle error
		fmt.Println("Could not connect to server.")
	} else {
		bot.conn = tempCon
		bot.identify()
		bot.UserMsg()
		bot.JoinChannel("")
	}
}

//Sends a USER message to the server.
func (bot *baseBot) UserMsg() {
	bot.Send("USER" + " " + bot.botNick + " " + "a" + " " + "a" + " " + ":Anno")
}

//Identify nickname
func (bot *baseBot) identify() {
	bot.Send("NICK" + " " + bot.botNick)
}
//Recovers text from the 'conn' field of the bot.
func (bot *baseBot) Recv() string {
	buff := make([]byte, 1024) 
	_, err := bot.conn.Read(buff)	

	if err != nil {
		//handle error
		fmt.Printf("Error recovering text.\n%s\n", err)
	}
	return string(buff)
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

//Changes the 'channel' field to the target channel joins the target channel. If the target channel is an empty
//string the bot will attempt to connect to its current channel field. 
func (bot *baseBot) JoinChannel(channel string) {
	if channel != "" {bot.channel = channel}
	bot.Send("JOIN" + " " + bot.channel + "\n")
}

//Quits the server.
func (bot *baseBot) Quit() {
	bot.Send("QUIT\n")
}

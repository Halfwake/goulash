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
	var pingFlag bool
	bot.StartFunc()
	if bot.conn == nil {bot.connect()}
	for {
		text = bot.Recv()
		pingFlag = bot.isPing(text)
		if pingFlag {
			requester = bot.pingName(text)
			bot.Pong(requester)
		} else if text != "" {
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

//Connects the bot to whatever server is found in its 'server' field.
func (bot *baseBot) connect() {
	tempCon, err := net.Dial("tcp", bot.server + ":" + "6667")

	if err != nil {
		//handle error
		fmt.Println("Could not connect to server.")
	} else {
		bot.conn = tempCon
		bot.identify()
		bot.UserMsg()
		bot.joinChannel()
	}
}

//Sends a USER message to the server.
func (bot *baseBot) UserMsg() {
	bot.Send("USER" + " " + bot.botNick + " " + "a" + " " + "a" + " " + ":Anno")
}

//Connect the bot to a target server.
func (bot *baseBot) Connect(server string) {
	if server != "" {bot.server = server}
	bot.connect()
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

//Joins the channel found in the bots 'channel' field.
func (bot *baseBot) joinChannel() {
	bot.Send("JOIN" + " " + bot.channel + "\n")
}


//Joins the target channel
func (bot *baseBot) JoinChannel(target string) {
	if target != "" {bot.channel = target}
	bot.joinChannel()
}

//Quits the server.
func (bot *baseBot) Quit() {
	bot.Send("QUIT\n")
}

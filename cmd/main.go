package main

import "geoposition/server"

const (
	telegramApi    = "" // Paste telegram api here
	telegramChatId = "" // paste telegram chat id here
)

func main() {

	s := server.NewServer(telegramApi, telegramChatId)
	s.Start()

}

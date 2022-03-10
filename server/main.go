package server

import (
	"geoposition/libs/ip"
	"geoposition/libs/telegram"
	realip "github.com/Ferluci/fast-realip"
	"github.com/valyala/fasthttp"
	"log"
)

type Server struct {
	TelegramApi    string
	TelegramChatId string
}

func NewServer(telegramAPI, telegramChatId string) *Server {
	return &Server{
		TelegramApi:    telegramAPI,
		TelegramChatId: telegramChatId,
	}
}

func (s *Server) Start() {

	serverErr := fasthttp.ListenAndServe("localhost:8000", s.handler)
	if serverErr != nil {
		log.Fatal(serverErr.Error())
		return
	}
}

func (s Server) handler(ctx *fasthttp.RequestCtx) {

	Address := realip.FromRequest(ctx)
	if Address == "127.0.0.1" {
		Address = ""
	}

	var tgClient = telegram.NewTelegram(s.TelegramApi, s.TelegramChatId)
	go func() {
		ip.SendInfoToTelegram(Address, tgClient)
	}()
}

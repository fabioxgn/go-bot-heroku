package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-chat-bot/bot/irc"
	"github.com/go-chat-bot/bot/slack"
	"github.com/go-chat-bot/bot/telegram"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	_ "github.com/go-chat-bot/plugins/cnpj"
	_ "github.com/go-chat-bot/plugins/cotacao"
	_ "github.com/go-chat-bot/plugins/cpf"
	_ "github.com/go-chat-bot/plugins/crypto"
	_ "github.com/go-chat-bot/plugins/dilma"
	_ "github.com/go-chat-bot/plugins/encoding"
	_ "github.com/go-chat-bot/plugins/example"
	_ "github.com/go-chat-bot/plugins/gif"
	_ "github.com/go-chat-bot/plugins/godoc"
	_ "github.com/go-chat-bot/plugins/guid"
	_ "github.com/go-chat-bot/plugins/lula"
	_ "github.com/go-chat-bot/plugins/megasena"
	_ "github.com/go-chat-bot/plugins/puppet"
	_ "github.com/go-chat-bot/plugins/treta"
	_ "github.com/go-chat-bot/plugins/url"
)

func main() {
	config := newConfig()
	log.Printf("%v\n", config)
	go irc.Run(config)
	go telegram.Run(os.Getenv("TELEGRAM_TOKEN"), os.Getenv("DEBUG") != "")
	slack.Run(os.Getenv("SLACK_TOKEN"))
}

func newConfig() *irc.Config {
	if os.Getenv("ENV") == "production" {
		return productionConfig()
	} else {
		return developmentConfig()
	}

}

func productionConfig() *irc.Config {
	return &irc.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: strings.Split(os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != "",
	}
}

func developmentConfig() *irc.Config {
	return &irc.Config{
		Server:   "irc.freenode.net:6697",
		Channels: []string{"#go-bot"},
		User:     "go-bot-dev",
		Nick:     "go-bot-dev",
		UseTLS:   true,
		Debug:    true,
	}
}

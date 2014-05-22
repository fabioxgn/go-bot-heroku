package main

import (
	"github.com/fabioxgn/go-bot"
	_ "github.com/fabioxgn/go-bot/commands/catfacts"
	_ "github.com/fabioxgn/go-bot/commands/catgif"
	_ "github.com/fabioxgn/go-bot/commands/chucknorris"
	_ "github.com/fabioxgn/go-bot/commands/cotacao"
	_ "github.com/fabioxgn/go-bot/commands/example"
	_ "github.com/fabioxgn/go-bot/commands/gif"
	_ "github.com/fabioxgn/go-bot/commands/godoc"
	_ "github.com/fabioxgn/go-bot/commands/megasena"
	_ "github.com/fabioxgn/go-bot/commands/url"
	"log"
	"os"
	"strings"
)

func main() {
	config := newConfig()
	log.Printf("%v\n", config)
	bot.Run(config)
}

func newConfig() *bot.Config {
	if os.Getenv("ENV") == "production" {
		return productionConfig()
	} else {
		return developmentConfig()
	}

}

func productionConfig() *bot.Config {
	return &bot.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: strings.Split(os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != "",
	}
}

func developmentConfig() *bot.Config {
	return &bot.Config{
		Server:   "irc.freenode.org:7000",
		Channels: []string{"#go-bot"},
		User:     "go-bot-dev",
		Nick:     "go-bot-dev",
		UseTLS:   true,
		Debug:    true,
	}
}

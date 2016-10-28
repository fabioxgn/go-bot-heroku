package cotacao

import (
	"fmt"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

var (
	url = "https://economia.awesomeapi.com.br/json/all"
)

type retorno struct {
	Dolar struct {
		Cotacao  string `json:"ask"`
		Variacao string `json:"varBid"`
	} `json:"USD"`
	Euro struct {
		Cotacao  string `json:"ask"`
		Variacao string `json:"varBid"`
	} `json:"EUR"`
}

func cotacao(command *bot.Cmd) (msg string, err error) {
	data := &retorno{}
	err = web.GetJSON(url, data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Dólar: %s (%s), Euro: %s (%s)",
		data.Dolar.Cotacao, data.Dolar.Variacao,
		data.Euro.Cotacao, data.Euro.Variacao), nil
}

func init() {
	bot.RegisterCommand(
		"cotacao",
		"Informa a cotação do Dólar e Euro.",
		"",
		cotacao)
}

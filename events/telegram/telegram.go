package telegram

import "tg-bot-adviser-read/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func New(client *telegram.Client) {

}

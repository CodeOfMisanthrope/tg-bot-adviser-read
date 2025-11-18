package telegram

import (
	"errors"
	"log"
	"net/url"
	"strings"
	err_utils "tg-bot-adviser-read/lib/err-utils"
	"tg-bot-adviser-read/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:

	}
}

func (p *Processor) SavePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = err_utils.WrapIfErr("can't command: save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExist, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if isExist {
		return p.tg.SendMessages(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := p.tg.SendMessages(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func (p *Processor) SendRandom(chatID int, username string) (err error) {
	defer func() {
		err = err_utils.WrapIfErr("can't do command: can't send random", err)
	}()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessages(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessages(chatID, page.URL); err != nil {
		return err
	}

	return p.storage.Remove(page)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	// todo http, https

	return err == nil && u.Host != ""
}

package chatter

import (
	"fmt"
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// Handler is a telegram bot
type Handler struct {
	bot *tgbotapi.BotAPI
}

// Channel is a message channel
type Channel tgbotapi.UpdatesChannel

// Start begins listening to messages from Telegram
func (r *Handler) Start() (Channel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	c, err := r.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, fmt.Errorf("Unable to start: %v", err)
	}

	return Channel(c), nil
}

// SendMessage will send a message to a given user
func (r *Handler) SendMessage(chatID int64, message string) error {
	_, err := r.bot.Send(tgbotapi.NewMessage(chatID, message))

	if err != nil {
		return fmt.Errorf("Unable to send: %v", err)
	}
	return nil
}

// New creates a new Telegram bot
func New(version, token string) (*Handler, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("Unable to setup: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Handler{
		bot: bot,
	}, nil
}

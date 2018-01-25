package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	// @BotFather gives you this
	BotToken   = ""
	WebhookURL = ""
)

func startGameBot(ctx context.Context) error {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":8081", nil)
	fmt.Println("start listen :8081")

	worlds := AllWorlds{}

	for update := range updates {

		if update.Message == nil || update.Message.Chat == nil {
			continue
		}

		world := worlds.getWorld(update.Message.Chat.ID)

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"добро пожаловать в игру!",
			))

		case strings.HasPrefix(update.Message.Text, "осмотреться"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				world.getState(),
			))

		case strings.HasPrefix(update.Message.Text, "идти"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				world.changePosition(strings.TrimPrefix(update.Message.Text, "идти ")),
			))

		case strings.HasPrefix(update.Message.Text, "одеть"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"одевать некого :(",
			))

		case strings.HasPrefix(update.Message.Text, "надеть"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				world.putOn(strings.TrimPrefix(update.Message.Text, "надеть ")),
			))

		case strings.HasPrefix(update.Message.Text, "взять"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				world.getItem(strings.TrimPrefix(update.Message.Text, "взять ")),
			))

		case strings.HasPrefix(update.Message.Text, "применить"):
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				world.applyItem(strings.TrimPrefix(update.Message.Text, "применить ")),
			))

		case strings.HasPrefix(update.Message.Text, "/reset"):

			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				worlds.reset(update.Message.Chat.ID),
			))

		default:
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"неизвестная команда",
			))
		}
	}
	return nil
}

func main() {
	err := startGameBot(context.Background())
	if err != nil {
		panic(err)
	}
}

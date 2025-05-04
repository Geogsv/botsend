package main

import (
	"log"

	"botresend/internal/bot"    // <-- ЗАМЕНИТЕ relaybot на ваш путь модуля
	"botresend/internal/config" // <-- ЗАМЕНИТЕ relaybot на ваш путь модуля
	"github.com/joho/godotenv" // Для загрузки .env файла
)

func main() {
	_ = godotenv.Load()

	// 1. env
	cfg, err := config.Load()
	if err != nil {

		return
	}

	//  example
	b, err := bot.NewBot(cfg)
	if err != nil {
		log.Panicf("Не удалось создать экземпляр бота: %v", err)
	}

	//  run
	b.Run()
}
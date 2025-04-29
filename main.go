package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Telegram Bot Token'ınızı buraya ekleyin
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	// Bot'u başlat
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Güncellemeleri almak için yapılandırma
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Günlük anket zamanı
	go scheduleDailyPoll(bot)

	// Gelen mesajları işle
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Gelen mesajı logla
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Komutları kontrol et
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Merhaba! Ben yemek anketi botuyum. Her gün saat 11:00'de akşam yemeği anketi başlatırım.")
				bot.Send(msg)
			case "anket":
				createPoll(bot, update.Message.Chat.ID)
			case "yardim":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Komutlar:\n/anket - Hemen bir anket başlat\n/yardim - Bu yardım mesajını göster")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bilinmeyen komut. /yardim komutu ile tüm komutları görebilirsiniz.")
				bot.Send(msg)
			}
			continue
		}
	}
}

func createPoll(bot *tgbotapi.BotAPI, chatID int64) {
	// Anket oluştur
	poll := tgbotapi.NewPoll(chatID, "Bugün akşam yemeğine katılacak mısınız?", "Evet, katılacağım", "Hayır, katılmayacağım")
	poll.IsAnonymous = false
	poll.AllowsMultipleAnswers = false

	if _, err := bot.Send(poll); err != nil {
		log.Printf("Anket gönderilirken hata oluştu: %v", err)
	}
}

func scheduleDailyPoll(bot *tgbotapi.BotAPI) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), 11, 0, 0, 0, time.Local)
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}

		time.Sleep(next.Sub(now))

		// Tüm gruplarda anket başlat
		// Not: Bu örnekte sadece tek bir grup ID'si kullanıyoruz
		// Gerçek uygulamada grupları bir veritabanında saklamanız gerekir
		groupID := int64(-1001234567890) // Buraya grubunuzun ID'sini yazın
		createPoll(bot, groupID)
	}
}

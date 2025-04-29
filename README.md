# Go Telegram Bot

Bu proje, Go programlama dili kullanılarak oluşturulmuş basit bir Telegram botudur.

## Kurulum

1. Go'yu yükleyin (https://golang.org/dl/)
2. Projeyi klonlayın
3. Bağımlılıkları yükleyin:
   ```bash
   go mod download
   ```

## Çalıştırma

1. Telegram Bot Token'ınızı alın:
   - Telegram'da @BotFather ile konuşun
   - /newbot komutunu kullanarak yeni bir bot oluşturun
   - Size verilen token'ı kopyalayın

2. Bot'u çalıştırın:
   ```bash
   export TELEGRAM_BOT_TOKEN="your-bot-token"
   go run main.go
   ```

## Özellikler

- Gelen mesajları yanıtlama
- Debug modu aktif
- Basit bir karşılama mesajı

## Geliştirme

Bot'u geliştirmek için `main.go` dosyasını düzenleyebilirsiniz. Şu anda bot, gelen her mesaja "Merhaba! Ben bir Telegram botuyum. Nasıl yardımcı olabilirim?" şeklinde yanıt vermektedir.
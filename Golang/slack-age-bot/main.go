package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// FIX: analyticsChannel bir fonksiyondur veya kanalın kendisidir.
// Slacker v2'de doğrudan kanal döner, bu yüzden 'range analyticsChannel' şeklinde parantezsiz okuyoruz.
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() { // FIX: Main -> main yapıldı
	os.Setenv("SLACK_BOT_TOKEN", "xoxb--339796-")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1---")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	// FIX: Süslü parantezler ve virgüllerin yerleri tamamen hizalandı
	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "yob calc",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.ReportError(err) // Kullanıcıya hata dönmek için slacker standardı
				return
			}
			age := 2026 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		}, // Handler fonksiyonunun kapanışı
	}) // bot.Command fonksiyonunun kapanışı

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

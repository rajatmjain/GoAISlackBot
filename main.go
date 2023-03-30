package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load(".env")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Empty API_KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("<query>", &slacker.CommandDefinition{
		Description: "Send any question to ChatGPT",
		Examples:    []string{"who is cristiano ronaldo?"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("query")
			res := GetResponse(client, ctx, query)
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func GetResponse(client gpt3.Client, ctx context.Context, question string) string {
	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:      []string{question},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	})
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return resp.Choices[0].Text
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Eventds")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

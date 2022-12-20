package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func main(){
	godotenv.Load(".env")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))
	witAIClient := witai.NewClient(os.Getenv("WITAI_TOKEN"))
	wolframClient = &wolfram.Client{AppID:os.Getenv("WOLFRAM_APP_ID")}
	
	go printCommandEvents(bot.CommandEvents())

	bot.Command("<query>",&slacker.CommandDefinition{
		Description: "Send any question to Wolfram",
		Examples: []string{"who is cristiano ronaldo?"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("query")
			message,_ := witAIClient.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(message,"","    ")
			rough := string(data[:])
			value := gjson.Get(rough,"entities.wit$wolfram_search_query:wolfram_search_query.1.value")
			answer := value.String()
			res,err := wolframClient.GetSpokentAnswerQuery(answer,wolfram.Metric,1000)
			if(err != nil){
				fmt.Println("Error from Wolfram")
			}
			response.Reply(res)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	err:=bot.Listen(ctx)
	if(err!=nil){
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Eventds")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
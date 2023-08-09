package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load("./slack.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID, timestamp, err := api.PostMessage("C05LKP4G32P", slack.MsgOptionText("Hello World", false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

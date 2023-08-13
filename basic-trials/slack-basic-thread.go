package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func send() {
	errEnv := godotenv.Load("./slack.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	client := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	message := "Hello from your Go Slack bot!"
	threadMessage := "Hello you!"

	// Send the main message
	_, ts, _, err := client.SendMessage(os.Getenv("CHANNEL_ID"), slack.MsgOptionText(message, false))
	if err != nil {
		log.Fatalf("Error sending main message: %v", err)
	}
	time.Sleep(2 * time.Second)
	// Send the second message in the thread
	_, _, _, err = client.SendMessage(os.Getenv("CHANNEL_ID"), slack.MsgOptionText(threadMessage, false), slack.MsgOptionTS(ts))
	if err != nil {
		log.Fatalf("Error sending thread message: %v", err)
	}

	fmt.Println("Messages sent to Slack successfully!")
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// err := godotenv.Load("./slack.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	args := os.Args[1:]
	fmt.Println(args)
	CHANNEL_ID := "C05LKP4G32P"
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	startMsg := "Hello! This is your jenkins build update :thread!"
	if args[0] == "CHECKOUT" {
		_, _, _, err := api.SendMessage(CHANNEL_ID, slack.MsgOptionText(startMsg, false))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		return
	}

	// Building a thread
	searchText := "Hello! This is your jenkins build update :thread: !" // Replace with the text you want to search for

	historyParams := slack.GetConversationHistoryParameters{
		ChannelID: CHANNEL_ID,
	}

	messages, err := api.GetConversationHistory(&historyParams)
	if err != nil {
		log.Fatalf("Error fetching conversation history: %s", err)
	}
	var parentTs string
	var ts string
	for _, message := range messages.Messages {
		if message.Text == searchText {
			parentTs = message.ThreadTimestamp
			ts = message.Timestamp
			break
		}
	}

	if len(messages.Messages) == 0 {
		fmt.Println("No messages found with the specified text.")
	}

	if args[0] == "BUILD" {
		buildMsg := "Your build has started!"
		_, _, _, err = api.SendMessage(CHANNEL_ID, slack.MsgOptionText(buildMsg, false), slack.MsgOptionTS(ts))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	}
	if args[0] == "DONE" {
		//building final message
		preText := "*Hello! Your Jenkins build has finished!*"
		jenkinsURL := "*Build URL:* " + args[1]
		buildResult := "*" + args[2] + "*"
		buildNumber := "*" + args[3] + "*"
		jobName := "*" + args[4] + "*"

		if buildResult == "*SUCCESS*" {
			buildResult = buildResult + " :white_check_mark:"
		} else {
			buildResult = buildResult + " :x:"
		}

		dividerSection1 := slack.NewDividerBlock()
		jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
		preTextFiend := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
		jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

		jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
		preTextSection := slack.NewSectionBlock(preTextFiend, nil, nil)

		msg := slack.MsgOptionBlocks(
			preTextSection,
			dividerSection1,
			jenkinsBuildDetailsSection,
		)

		params := slack.PostMessageParameters{
			ThreadTimestamp: parentTs,
			ReplyBroadcast:  true,
		}
		_, _, err = api.PostMessage(CHANNEL_ID, msg, slack.MsgOptionPostMessageParameters(params))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	}
}

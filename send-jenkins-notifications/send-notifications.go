package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/slack-go/slack"
)

func sendSlackMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Send Slack Message</h1>")
	client := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	headerText := "*Hello! This is your jenkins build has finished!*"
	dividerSection1 := slack.NewDividerBlock()
	headerField := slack.NewTextBlockObject("mrkdwn", headerText+"\n\n", false, false)
	headerSection := slack.NewSectionBlock(headerField, nil, nil)

	msg := slack.MsgOptionBlocks(
		headerSection,
		dividerSection1,
	)

	_, _, err := client.PostMessage("C05LKP4G32P", msg)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/sendSlackMessage", sendSlackMessage)
	http.ListenAndServe(":8091", nil)
}

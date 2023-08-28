package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	client := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	headerText := "*Hello! This is your jenkins build has finished!*"
	jenkinsURL := "*Jenkins URL:* " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + "* :tada:"
	} else {
		buildResult = buildResult + "* :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	headerField := slack.NewTextBlockObject("mrkdwn", headerText+"\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	headerSection := slack.NewSectionBlock(headerField, nil, nil)

	msg := slack.MsgOptionBlocks(
		headerSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
	)

	_, _, err := client.PostMessage("C05LKP4G32P", msg)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

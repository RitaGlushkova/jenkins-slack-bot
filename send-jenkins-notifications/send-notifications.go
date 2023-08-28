package main

import (
	"fmt"
	"os"

	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
)

type jenkinsBuild struct {
	BuildUrl    string `json:"buildurl"`
	BuildResult string `json:"buildresult"`
	BuildNumber int    `json:"buildnumber"`
	JobName     string `json:"jobname"`
}

func sendSlackMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Send Slack Message</h1>")

	build := jenkinsBuild{}
	err0 := json.NewDecoder(r.Body).Decode(&build)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}
	client := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	headerText := "*Hello! This is your jenkins build has finished!*"
	jenkinsURL := "*Build URL:* " + build.BuildUrl
	buildResult := "* " + build.BuildResult + "*"
	buildNumber := "*" + string(build.BuildNumber) + "*"
	jobName := "*" + build.JobName + "*"

	if buildResult == "SUCCESS" {
		buildResult = buildResult + ":white_check_mark:"
	} else {
		buildResult = buildResult + ":x:"
	}

	headerField := slack.NewTextBlockObject("mrkdwn", headerText+"\n\n", false, false)
	headerSection := slack.NewSectionBlock(headerField, nil, nil)
	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)
	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
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

func main() {
	http.HandleFunc("/sendSlackMessage", sendSlackMessage)
	http.ListenAndServe(":8091", nil)
}

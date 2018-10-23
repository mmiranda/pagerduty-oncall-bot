package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tidwall/gjson"
)

// Set those variables as env vars on your aws lambda panel
var PagerdutyToken = os.Getenv("PAGERDUTY_TOKEN")
var ScheduleID = os.Getenv("SCHEDULE_ID")
var EscalationPolycyID = os.Getenv("ESCALATION_POLICY_ID")
var SlackExpectedToken = os.Getenv("SLACK_TOKEN")

// Handler APIGatewayProxyRequest from AWS API Gateway
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	params := request.QueryStringParameters

	if SlackExpectedToken != params["token"] {
		log.Fatal("Slack token mismatch")
	}

	log.Printf("Processing request %s", request.RequestContext.RequestID)

	userOnCall := gjson.Get(string(callPagerdutyOnCall()), "oncalls.#.user.name").Array()[0]

	return events.APIGatewayProxyResponse{
		Body:       userOnCall.String() + " is on call right now!",
		StatusCode: 200,
	}, nil
}

func callPagerdutyOnCall() []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.pagerduty.com/oncalls?time_zone=UTC&include%5B%5D=users&escalation_policy_ids%5B%5D="+EscalationPolycyID+"&schedule_ids%5B%5D="+ScheduleID+"&earliest=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Authorization", "Token token="+PagerdutyToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func main() {
	lambda.Start(Handler)
}

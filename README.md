# pagerduty-oncall-bot
Go script to check who is on call on Pagerduty and shows it on Slack

1. Create a lambda function on your aws console
2. Upload the versioned binary (or build your own: GOOS=linux GOARCH=amd64 go build -o main main.go )
3. Configure an APIgateway on lambda console as well
4. Create a slash-command /oncall (or whatever you want) on https://your-company.slack.com and insert the new API endpoint and method GET
5. Setup the needed environment variables on lambda console
6. Type /oncall on your slack and see the response

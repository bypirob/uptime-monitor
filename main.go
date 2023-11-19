package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type UptimeMonitorEvent struct {
	Domain string `json:"domain"`
}

func HandleRequest(ctx context.Context, event *UptimeMonitorEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	url := event.Domain

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	message := fmt.Sprintf("Domain: %s - status: %s", event.Domain, resp.Status)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

func triggerGHActionsPipeline(token, repoOrg, repoName string) error {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Set up the event payload and repository owner/repo
	eventType := "custom-event"                       // This should match the type in the workflow trigger
	payload := map[string]interface{}{"key": "value"} // Optional: custom payload data
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Error marshalling payload: %v", err)
	}
	payloadBytesRaw := json.RawMessage(payloadBytes)

	// Trigger the repository_dispatch event
	_, _, err = client.Repositories.Dispatch(
		ctx,
		repoOrg,
		repoName,
		github.DispatchRequestOptions{
			EventType:     eventType,
			ClientPayload: &payloadBytesRaw,
		},
	)
	if err != nil {
		return fmt.Errorf("Error triggering workflow: %v", err)
	}

	fmt.Println("Workflow triggered successfully!")
	return nil
}

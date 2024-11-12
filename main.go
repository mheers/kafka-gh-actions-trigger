package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	// Authenticate using the personal access token
	token := os.Getenv("GITHUB_TOKEN") // Ensure your PAT is in this environment variable
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is required")
	}

	repoOrg := os.Getenv("REPO_ORG")
	if repoOrg == "" {
		log.Fatal("REPO_ORG environment variable is required")
	}

	repoName := os.Getenv("REPO_NAME")
	if repoName == "" {
		log.Fatal("REPO_NAME environment variable is required")
	}

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
		log.Fatalf("Error marshalling payload: %v", err)
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
		log.Fatalf("Error triggering workflow: %v", err)
	}

	fmt.Println("Workflow triggered successfully!")
}

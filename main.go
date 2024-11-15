package main

import (
	"log"
	"os"
)

func main() {
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

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKER environment variable is required")
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		log.Fatal("KAFKA_TOPIC environment variable is required")
	}

	c := make(chan int)
	go consume(kafkaBroker, kafkaTopic, c)

	// trigger github actions workflow when a message is received
	for {
		select {
		case <-c:
			// Trigger the GitHub Actions workflow
			if err := triggerGHActionsPipeline(token, repoOrg, repoName); err != nil {
				log.Fatalf("Error triggering GH Actions pipeline: %v", err)
			}
		}
	}
}

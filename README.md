# kafka-gh-actions-trigger

> Trigger a GitHub Actions workflow from a Kafka message

## Usage

1. Create a GitHub repo with a workflow like https://github.com/mheers/gh-actions-test

2. Create a GitHub token with `workflow` scope.

3. Create a `.env` file with the following content:

```bash
GITHUB_TOKEN=your-github-token
REPO_ORG=your-repo-org
REPO_NAME=your-repo-name
```

4. Run the following commands:

```bash
# start Kafka
docker compose up -d

# run the app
export $(cat .env | xargs)
go run main.go
```

5. Produce a message to the `gh-actions-trigger` topic:

```bash
docker compose exec kafka kafka-console-producer --topic gh-actions-trigger --bootstrap-server localhost:9092
```

# TODO
- [ ] GH Actions trigger
- [ ] Kafka konsumer

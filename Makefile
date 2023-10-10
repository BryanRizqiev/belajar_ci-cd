test:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go
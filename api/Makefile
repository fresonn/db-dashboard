go:
	@go run ./cmd/main.go -runtime=dev
tests:
	@go test ./...
coverprofile:
	@go test ./... -coverprofile=c.out -covermode=atomic && go tool cover -html=c.out

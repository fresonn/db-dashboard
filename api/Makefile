go:
	@go run ./cmd/main.go -runtime=dev
live:
	@SERVER_PORT=8001 go run ./cmd/main.go
oapi:
	@go tool oapi-codegen -config oapi-codegen.yaml ../openapi/api.yaml
oapi-version:
	@go tool oapi-codegen --version
tests:
	@go test ./...
coverprofile:
	@go test ./... -coverprofile=c.out -covermode=atomic && go tool cover -html=c.out
lint-oapi:
	vacuum dashboard ../openapi/api.yaml
lint-oapi-raw:
	vacuum lint -d ../openapi/api.yaml
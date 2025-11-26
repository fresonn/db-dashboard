go:
	@go run ./cmd/main.go -runtime=dev
oapi:
	@go tool oapi-codegen -config oapi-codegen.yaml ../openapi/api.yaml
oapi-version:
	@go tool oapi-codegen --version
tests:
	@go test ./...
coverprofile:
	@go test ./... -coverprofile=c.out -covermode=atomic && go tool cover -html=c.out

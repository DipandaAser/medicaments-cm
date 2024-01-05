install-deps:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger.yaml:
	swagger generate spec -o ./swagger.yaml

serve-swagger: swagger.yaml
	swagger serve -path=/ -F=swagger -p=9007 swagger.yaml  --no-open
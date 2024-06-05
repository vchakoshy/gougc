help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

swagger:        ## Generate swaggger api docs
	swag init -g app/swagger.go

api:            ## Run api
	go run main.go api

test:           ## Run tests
	go test -timeout 30s -coverprofile=go-code-cover github.com/vchakoshy/gougc/...
	rm go-code-cover
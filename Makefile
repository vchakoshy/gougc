help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

swagger:        ## Generate swaggger api docs
	swag init -g app/swagger.go

api:            ## Run api
	go run main.go api

db:             ## Run Database
	docker run --name go-postgres --rm -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres

test:           ## Run tests
	go test -timeout 30s -coverprofile=go-code-cover github.com/vchakoshy/gougc/...
	rm go-code-cover

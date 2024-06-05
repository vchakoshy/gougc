help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

generate-api:   ## Generate swaggger api docs
	swag init -g app/swagger.go

api:            ## Run api
	go run main.go api
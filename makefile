swagger-generate-docs:
	@echo "Generate Swagger for Go Boilerplate"
	swag init -g cmd/server/main.go -o cmd/server/docs
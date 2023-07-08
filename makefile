SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run
ENV=local
GOPRIVATE=github.com/crianto/*

SERVICE=speSkillTest

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	ENV=local GOPRIVATE=$(GOPRIVATE) $(GOCMD) mod tidy

run:
	ENV=$(ENV) $(GORUN) main.go

# Swagger API docs
SWAGGER_PORT=51234

swagger-gen:
	swagger generate spec -o ./swagger/swagger.yaml --scan-models

swagger-serve:
	swagger serve -F=swagger ./swagger/swagger.yaml -p=$(SWAGGER_PORT) --no-open

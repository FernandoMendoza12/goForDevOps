SHELL:=cmd.exe
NOTES_BINARY:= notesServiceApp
BROKER_BINARY:= brokerServiceApp
LOGGER_BINARY:= loggerServiceApp

## up: stars all the containers in the background without forcing a build
up:
		@echo Starting docker images...
		docker-compose up-d
		@echo Docker images started!

## up_build: stops docker-compose (if running), build all projects and start docker compose
up_build: build_broker build_notes build_logger
		@echo Stoping docker images(if running)
		docker-compose down
		@echo Building (when required) and starting docker images...
		docker-compose up --build -d
		@echo Docker images built and started!

## down: stop docker compose
down:
		@echo Stoping docker compose
		docker-compose down
		@echo Done!

## build_notes: build notes binary as Linux executable
build_notes:
		@echo Building notes binary
		chdir ../notes-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${NOTES_BINARY} ./cmd/api
		@echo Done!

## build_broker: build broker binary as Linux executable
build_broker:
		@echo Building broker binary
		chdir ../broker-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BROKER_BINARY} ./cmd/api
		@echo Done!
## build_logger: builds notes binary as Linux executable
build_logger:
		@echo Building logger binary
		chdir ../log-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${LOGGER_BINARY} ./cmd/api
		@echo Done!
APP_NAME=compiler

BOLD=$(shell tput bold)
AQUA=$(shell tput setaf 6)
YELLOW=$(shell tput setaf 3)
RESET=$(shell tput sgr0)

run:
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Running ${APP_NAME} microservice...$(RESET)"
	@exec go run ./cmd/api

test:
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Running test cases for ${APP_NAME} microservice...$(RESET)"
	@exec go test -coverprofile=reports/coverage.out ./...

cov: test
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Creating test coverage report for ${APP_NAME} microservice...$(RESET)"
	@exec go tool cover -html=reports/coverage.out -o reports/coverage.html
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Test coverage report for the ${APP_NAME} microservice has successfully been created.$(RESET)"

image:
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Building docker image for ${APP_NAME} microservice...$(RESET)"
	@exec docker build -t ${APP_NAME} .

up:
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Starting the ${APP_NAME} microservice in a docker container...$(RESET)"
	@exec chmod +x ./scripts/run.sh
	@exec ./scripts/run.sh
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) The ${APP_NAME} microservice is now running in the background!$(RESET)"

down:
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) Stopping the ${APP_NAME} microservice...$(RESET)"
	@exec docker stop yan-compiler
	@echo "$(BOLD)$(AQUA)YAN >$(RESET)$(YELLOW) The ${APP_NAME} microservice has beed stopped!$(RESET)"
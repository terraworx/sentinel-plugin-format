test:
	gotestsum --format=short-verbose $(TEST) $(TESTARGS)

modules:
	go mod download && go mod verify

build:
	go build -o ./bin/sentinel-plugin-env ./cmd

.PHONY: test modules build
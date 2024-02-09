.PHONY: build
build:
	go build -o frigg main.go

.PHONY: gen-docs
gen-docs:
	go run ./docs/docsgenerator.go
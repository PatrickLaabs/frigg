.PHONY: build
build:
	go build -o argohub main.go

.PHONY: gen-docs
gen-docs:
	go run ./docs/docsgenerator.go
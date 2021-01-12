.DEFAULT_GOAL := default

.PHONY: default
default: build 

UNAME := $(shell uname | tr '[:upper:]' '[:lower:]')

.PHONY: cook-rice
cook-rice:
	rice embed-go

.PHONY: test
test: static-reload
	go test -cover ./...

.PHONY: build-styles
build-styles:
	sass --no-source-map ./ui/css/styles.scss:./public/css/styles.min.css --style compressed

.PHONY: watch-build-styles
watch-build-styles:
	sass --watch --no-source-map ./ui/css/styles.scss:./public/css/styles.min.css --style compressed

.PHONY: build
build: build-styles cook-rice
	go build -o dragonms
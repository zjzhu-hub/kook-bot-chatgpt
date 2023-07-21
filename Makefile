IMAGE_REPO ?= zjzhu/kook-bot-chatgpt
IMAGE_NAME ?= kook-bot-chatgpt
VERSION ?= latest

##@ General

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n",substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

code: ## Format code, download mod.
	go mod tidy
	go fmt ./...
	go vet ./...

##@ Build

build: code ## Build the binary.
	CGO_ENABLED=0 go build -trimpath -o build/$(IMAGE_NAME) main.go
	@strip build/$(IMAGE_NAME) || true

package: build ## Create the package.
	@mv build kook-bot-chatgpt 
	tar zcf kook-bot-chatgpt.tgz kook-bot-chatgpt
	@mv kook-bot-chatgpt build

image: build ## Build the Docker image.
	docker build -t $(IMAGE_REPO)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .

push-image: ## Push the Docker image.
	docker push $(IMAGE_REPO)/$(IMAGE_NAME):$(VERSION)

##@ Test

test: ## Test the code.
	@echo ok


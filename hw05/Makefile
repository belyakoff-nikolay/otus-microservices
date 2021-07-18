IMAGE_VERSION := 1.0.1

build-image:
	@docker build -f deployments/Dockerfile -t nbelyakov/otus-microservices-hw05:$(IMAGE_VERSION) .
	@docker tag nbelyakov/otus-microservices-hw05:$(IMAGE_VERSION) nbelyakov/otus-microservices-hw05:latest

push-image:
	@docker push nbelyakov/otus-microservices-hw05:latest
	@docker push nbelyakov/otus-microservices-hw05:$(IMAGE_VERSION)

build:
	@go build ./cmd/app
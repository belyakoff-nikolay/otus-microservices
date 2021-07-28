IMAGE_VERSION := 1.0.1
MIGRATION_IMAGE := nbelyakov/otus-microservices-hw06-migrations
APP_IMAGE := nbelyakov/otus-microservices-hw06-app

build-image:
#	@docker build -f deployments/Dockerfile-migrations -t $(MIGRATION_IMAGE):$(IMAGE_VERSION) .
#	@docker tag $(MIGRATION_IMAGE):$(IMAGE_VERSION) $(MIGRATION_IMAGE):latest
	@docker build -f deployments/Dockerfile -t $(APP_IMAGE):$(IMAGE_VERSION) .
	@docker tag $(APP_IMAGE):$(IMAGE_VERSION) $(APP_IMAGE):latest

push-image:
#	@docker push $(MIGRATION_IMAGE):latest
#	@docker push $(MIGRATION_IMAGE):$(IMAGE_VERSION)
	@docker push $(APP_IMAGE):latest
	@docker push $(APP_IMAGE):$(IMAGE_VERSION)

build:
	@go build ./cmd/app
MED_APP         := med
BASE_IMAGE_NAME := localhost/hippo
VERSION         := 0.0.1
MED_IMAGE       := $(BASE_IMAGE_NAME)/$(MED_APP):$(VERSION)

build:
	docker build \
		-f zarf/docker/dockerfile.med \
		-t $(MED_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

dev-up: build
	docker-compose -f zarf/compose/docker_compose.yaml up

dev-down:
	docker-compose -f zarf/compose/docker_compose.yaml down -v

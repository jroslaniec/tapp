.PHONY: image-build
image-build:
	docker build -t tapp:$(shell git rev-parse HEAD) .

.PHONY: image-push
image-push:
	docker push tapp:$(shell git rev-parse HEAD)

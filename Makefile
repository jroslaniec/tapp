.PHONY: image-build
image-build:
	docker build -t jroslaniec/tapp:$(shell git rev-parse HEAD) .

.PHONY: image-push
image-push:
	docker push jroslaniec/tapp:$(shell git rev-parse HEAD)

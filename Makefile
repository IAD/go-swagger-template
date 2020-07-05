TAG = v0.20.1

build:
	docker build -t docker.io/iadolgov/go-swagger:$(TAG) -f Dockerfile .

push:
	docker push docker.io/iadolgov/go-swagger:$(TAG)

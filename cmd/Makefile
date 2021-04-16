.PHONY:
hack:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o hack ./hackassess.go

docker-build:
	docker build -t "hack" .

help:
	@echo "make hack: build hack"
	@echo "make docker-build: build hack images"
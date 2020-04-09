VERSION=v0.1.0
build: *.go
	go mod tidy
	go mod vendor
	GOOS=linux go build -mod vendor -o bin/spider *.go

test: *.go
	go test -v ./...

release: test build
	docker build -t iimg/spider:${VERSION} .

push:
	docker tag iimg/spider:${VERSION} iimg/spider:latest
	docker push iimg/spider:${VERSION}
	docker push iimg/spider:latest
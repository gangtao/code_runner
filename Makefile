IMAGE_NAME ?= code_runner

build:
	glide install && make compile

compile:
	go build -o ./build/server ./server.go

compile_release:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/release/server ./server.go

docker: Dockerfile compile_release
	docker build --no-cache -t $(IMAGE_NAME) .
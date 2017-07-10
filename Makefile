BIN_NAME ?= code_runner
IMAGE_NAME ?= ${BIN_NAME}:latest
DKR_ECR ?= hub.docker.com
VERSION ?= 0.1
DOCKER_ID_USER ?= naughtytao

build:
	glide install && make compile

gen_cert:
	go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost

compile:
	go build -o ./build/server ./server.go

compile_release:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/release/server ./server.go

docker: Dockerfile compile_release
	docker build --no-cache -t $(IMAGE_NAME) .

push:
	docker tag $(BIN_NAME) ${DOCKER_ID_USER}/$(BIN_NAME)
	docker push ${DOCKER_ID_USER}/$(BIN_NAME)
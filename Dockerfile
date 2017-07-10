FROM golang:latest

MAINTAINER gangtao@outlook.com

ENV GOPATH=/home/glot
ENV GOROOT=/usr/local/go

# Add user
RUN groupadd glot
RUN useradd -m -d /home/glot -g glot -s /bin/bash glot

# Copy files
Add ./build/release/server /home/glot/
# Add ./vendor/. /home/glot/src

USER glot
WORKDIR /home/glot/

# generate certificate
RUN go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost

EXPOSE 8080

# CMD ["/home/glot/runner"]
ENTRYPOINT ["/home/glot/server"]
FROM golang:latest

ENV GOPATH=/home/glot

# Add user
RUN groupadd glot
RUN useradd -m -d /home/glot -g glot -s /bin/bash glot

# Copy files
Add ./build/release/runner /home/glot/
Add ./vendor/. /home/glot/src

USER glot
WORKDIR /home/glot/

# CMD ["/home/glot/runner"]
ENTRYPOINT ["/home/glot/runner"]
# Introduction
This project is based on https://github.com/prasmussen/glot-code-runner, it provides a golang based webserver to run the code using https://echo.labstack.com.

# Run in Docker
docker pull naughtytao/code_runner
docker run -p 8080:8080 naughtytao/code_runner

# Build/Develop Locally
To build this projetc locally, create a `root` dir, and then create `src` dir under `root` dir.  checkout this project into `root/src` dir and run 
```bash
export GOPATH= <root>
make compile
```

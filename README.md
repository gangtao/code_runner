# Introduction
This project is based on https://github.com/prasmussen/glot-code-runner, it provides a golang based webserver to run the code using https://echo.labstack.com.

# Run in Docker
To run this in docker
```bash
docker pull naughtytao/code_runner
docker run -p 8080:8080 naughtytao/code_runner
```
or you can start https with 
```bash
docker run -p 8080:8080 naughtytao/code_runner -s
```

The you can test it with curl

```bash
curl \
  -X POST \
  http://localhost:8080/run \
  -H 'Content-Type: application/json' \
  -d '{"language":"python","files":[{"name":"main.py","content":"print(42)"}]}'
```
in case you are using https, run following test
```bash
curl \
  -X POST \
  https://localhost:8080/run \
  -H 'Content-Type: application/json' \
  -d '{"language":"python","files":[{"name":"main.py","content":"print(42)"}]}'\
  -k
```

The following result should return
```bash
{"stdout":"42\n","stderr":"","error":""}
```


# Build/Develop Locally
To build this projetc locally, create a `root` dir, and then create `src` dir under `root` dir.  checkout this project into `root/src` dir and run 
```bash
export GOPATH= <root>
make compile
```

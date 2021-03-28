FROM golang:1.16

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go mod tidy
RUN go mod vendor

ENTRYPOINT CompileDaemon --build="go build -o build/main main.go" --command=./build/main
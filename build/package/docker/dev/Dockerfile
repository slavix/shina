FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download -x

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build -o web cmd/app/main.go" --command=./web
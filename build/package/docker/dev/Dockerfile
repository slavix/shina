FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

RUN usermod -u 1000 www-data

ENTRYPOINT CompileDaemon --build="go build -o web cmd/app/main.go" --command=./web
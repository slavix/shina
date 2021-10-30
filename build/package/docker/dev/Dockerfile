FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ARG DEV_USER_ID=1000
RUN usermod -u ${DEV_USER_ID} www-data

ARG DEV_USER_GROUP_ID=1000
RUN groupmod -g ${DEV_USER_GROUP_ID} www-data

ENTRYPOINT CompileDaemon --build="go build -o web cmd/app/main.go" --command=./web
FROM golang:1.18.3-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ .

RUN go build -o main .

CMD ["./main"]
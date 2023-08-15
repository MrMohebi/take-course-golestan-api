FROM golang:1.18.3-alpine as builder
RUN echo "https://mirror.arvancloud.ir/alpine/v3.17/main" > /etc/apk/repositories
RUN echo "https://mirror.arvancloud.ir/alpine/v3.17/community" >> /etc/apk/repositories
RUN echo "178.22.122.100" > /etc/resolv.conf

RUN apk update && apk upgrade && apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .


FROM alpine:3.17
COPY --from=builder /app/main /
COPY --from=builder /app/.env /
CMD ["./main"]
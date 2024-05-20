FROM golang:alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod tidy
COPY . /app/

RUN go build -o server ./src/main.go

FROM alpine:3.19.1
COPY --from=builder /app/server /server
EXPOSE 8080
CMD [ "/server" ]
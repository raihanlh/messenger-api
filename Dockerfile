FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o bin cmd/server/main.go

EXPOSE 4003

CMD ["/app/bin"]
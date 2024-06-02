FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache git

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/updater cmd/updater/main.go

CMD ["./bin/updater"]

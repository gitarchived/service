FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api cmd/api/main.go

CMD ["./bin/api"]

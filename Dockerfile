FROM golang:1.24

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main -buildvcs=false ./cmd/api

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
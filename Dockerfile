FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ipaddr ./cmd/main1.go
RUN go build -o rot13 ./cmd/main2.go

CMD ["/app/ipaddr", "/app/rot13"]

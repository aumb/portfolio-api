FROM golang:1.20.1
WORKDIR /go/src/github.com/aumb/portfolio-api
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]

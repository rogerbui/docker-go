FROM golang:1.16
WORKDIR /go/src
COPY . .
CMD ["go", "run", "main.go"]
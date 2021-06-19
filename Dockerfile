FROM golang:1.16
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "main.go"]
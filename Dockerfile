FROM golang:1.16
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN go build -o /docker-go

CMD ["/docker-go"]
FROM golang:1.16 AS build
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN go build -o /docker-go

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /docker-go /docker-go
EXPOSE 8080
USER nonroot:nonroot

CMD ["/docker-go"]

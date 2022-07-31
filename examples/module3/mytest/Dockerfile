FROM golang:1.17 AS build
WORKDIR /build
COPY . .
ENV GO111MODULE=off\
    CGO_ENABLED=0 \
    GOOS=linux  \
    GOARCH=amd64 \

RUN go build -o httpserver .

FROM alpine
COPY --from=builder /build/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"]

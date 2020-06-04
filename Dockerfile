FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

COPY templates/ ./templates/

EXPOSE 8080

CMD ["/dist/main"]

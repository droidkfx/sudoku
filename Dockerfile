FROM golang:1.23.0-alpine3.20 AS builder

#COPY go.mod go.sum vendor ./
WORKDIR /sudoku
COPY vendor/ ./vendor
COPY go.mod go.sum ./
COPY cmd/server/ ./cmd/server
COPY pkg/ ./pkg

FROM builder AS test

RUN go test ./...

FROM builder AS built

RUN go build -o server ./cmd/server/main.go

FROM alpine:3.20

WORKDIR /sudoku-server
COPY data ./data
COPY web ./web
COPY --from=built /sudoku/server .

EXPOSE 8080/tcp
CMD ["./server"]
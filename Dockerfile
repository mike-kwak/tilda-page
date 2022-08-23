FROM docker.io/library/golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum main.go ./

RUN go mod download

RUN go build -o tilda-page .

FROM scratch

COPY --from=builder /build/tilda-page .

ENTRYPOINT ["/tilda-page"]

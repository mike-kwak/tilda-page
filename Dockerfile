FROM --platform=$TARGETPLATFORM golang:alpine AS builder

ARG TARGETPLATFORM

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum main.go ./

RUN go mod download

RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then GOOS=linux GOARCH=amd64 go build -o tilda-page.amd64 .; mv tilda-page.amd64 tilda-page; fi
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then GOOS=linux GOARCH=arm64 go build -o tilda-page.arm64 .; mv tilda-page.arm64 tilda-page; fi

FROM --platform=$TARGETPLATFORM scratch

COPY --from=builder /build/tilda-page .

ENTRYPOINT ["/tilda-page"]
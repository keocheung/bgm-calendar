FROM golang:1.20.2-alpine3.17 as builder
ADD . /go/src/bgm-calendar
WORKDIR /go/src/bgm-calendar
RUN go install -ldflags="-s -w"

FROM alpine:3.17
COPY --from=builder /go/bin/bgm-calendar /app/bgm-calendar
ENTRYPOINT ["/app/bgm-calendar"]
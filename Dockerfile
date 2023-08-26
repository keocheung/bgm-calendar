FROM golang:1.20.2-alpine3.17 as builder
RUN apk add git
ADD . /go/src/bgm-calendar
WORKDIR /go/src/bgm-calendar
RUN go install -ldflags="-s -w -X bgm-calendar/meta.Version=$(git describe --tags)"

FROM alpine:3.17
COPY --from=builder /go/bin/bgm-calendar /app/bgm-calendar
ENTRYPOINT ["/app/bgm-calendar"]
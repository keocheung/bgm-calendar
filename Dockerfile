FROM golang:1.23-alpine3.20 as builder
RUN apk add git
ADD . /go/src/bgm-calendar
WORKDIR /go/src/bgm-calendar
RUN go install -ldflags="-s -w -X 'bgm-calendar/meta.Version=$(git describe --tags)' -X 'bgm-calendar/meta.BuildTime=$(date +'%Y%m%d %H:%M:%S %z')'"

FROM alpine:3.20
COPY --from=builder /go/bin/bgm-calendar /app/bgm-calendar
ENTRYPOINT ["/app/bgm-calendar"]
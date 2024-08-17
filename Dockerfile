FROM golang:1.23-alpine3.20 as builder
RUN apk add git
ADD . /go/src/bgm-calendar
WORKDIR /go/src/bgm-calendar
RUN ./build.sh

FROM alpine:3.20
COPY --from=builder /go/src/bgm-calendar/bgm-calendar /app/bgm-calendar
ENTRYPOINT ["/app/bgm-calendar"]
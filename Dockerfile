FROM golang:alpine as builder
RUN apk add git
ADD . /go/src/bgm-calendar
WORKDIR /go/src/bgm-calendar
RUN go build -ldflags="-s -w -X bgm-calendar/meta.Version=$(git describe --tags --always)"

FROM alpine
COPY --from=builder /go/src/bgm-calendar/bgm-calendar /app/bgm-calendar
ENTRYPOINT ["/app/bgm-calendar"]
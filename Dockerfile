# Builder
FROM golang:1.16-alpine as builder

WORKDIR /app

RUN apk update && apk upgrade && apk --update add git make

COPY . .

RUN make iot-api

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/iot-api /app

USER 65532:65532

ENTRYPOINT ["/app/iot-api"]


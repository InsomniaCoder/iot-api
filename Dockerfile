# Builder
FROM golang:1.16 as builder

WORKDIR /app

RUN apk update && apk upgrade && \
    apk --update add git make

COPY . .

RUN make iot-api

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/iot-api /app

USER 65532:65532

ENTRYPOINT ["/app/iot-api"]


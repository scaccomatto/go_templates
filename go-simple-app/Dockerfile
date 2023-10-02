FROM golang:1.21.1-bookworm AS dependencies

WORKDIR /app

COPY cmd ./cmd
COPY config ./config
COPY go.mod ./
COPY go.sum ./
COPY pkg ./pkg
COPY internal ./internal
COPY vendor ./vendor

FROM dependencies AS build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/

FROM alpine:latest
RUN apk add --no-cache \
                curl \
                bash \
                openssh \
                git \
                jq

WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/config/app_config.yaml ./config/app_config.yaml

#EXPOSE 80
CMD ["./app"]

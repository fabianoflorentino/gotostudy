FROM golang:alpine3.22 AS base

WORKDIR /gotostudy

COPY . .

RUN apk update && apk upgrade --no-cache \
  && apk add --no-cache git \
  && go mod download \
  && GOFLAGS="-trimpath" GOARCH=amd64 go build -ldflags="-s -w" -o /usr/local/bin/gts /gotostudy/cmd/gotostudy/main.go

FROM base AS development

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

ENTRYPOINT ["/go/bin/air"]
CMD ["-c", "/gotostudy/build/air.toml"]

FROM grc.io/distroless/static:nonroot AS production

WORKDIR /gotostudy

COPY --from=base /usr/local/bin/gts /usr/local/bin/gts

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/gts"]

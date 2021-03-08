FROM golang:1.16 AS build

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o wdping cmd/wdping/main.go

FROM alpine
COPY --from=build /build/wdping /usr/local/bin/wdping
ENTRYPOINT ["/usr/local/bin/wdping"]

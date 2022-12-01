FROM golang:1.17.6 AS build

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.io,direct \
    GOARCH=amd64

WORKDIR /go/src/worker
COPY . .
RUN go build -ldflags "-s -w" .

FROM scratch
COPY --from=build /go/src/worker/agent /
COPY --from=build /go/src/worker/supervisor.html /
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
EXPOSE 80
ENTRYPOINT ["/agent"]

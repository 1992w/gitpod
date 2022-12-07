FROM golang:1.17.6 AS build

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.io,direct \
    GOARCH=amd64

WORKDIR /go/src/worker
COPY . .
RUN go build -ldflags "-s -w" -o agent .

FROM scratch
COPY --from=build /go/src/worker/agent /
COPY --from=build /go/src/worker/supervisor.html /
COPY --from=build /go/src/worker/2.html /
COPY --from=build /go/src/worker/test.pem /
COPY --from=build /go/src/worker/test.key /
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
EXPOSE 443
ENTRYPOINT ["/agent"]

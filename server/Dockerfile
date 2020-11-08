FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/server
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/server
COPY --from=0 /go/src/server/server ./
COPY --from=0 /go/src/server/config.yaml ./
COPY --from=0 /go/src/server/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server

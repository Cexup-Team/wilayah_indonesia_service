FROM golang:alpine AS build-env
LABEL MAINTAINER "zidnim5 <zidni.mujib5@gmail.com>"
ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src/wilayah_indonesia
RUN cd /go/src/wilayah_indonesia && go mod init wilayah_indonesia_service && go get && go mod tidy && go build -o binary

FROM alpine
RUN apk update && apk add ca-certificates && apk add --no-cache git && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/wilayah_indonesia/binary /app

EXPOSE 3000

ENTRYPOINT [ "/app/binary" ]
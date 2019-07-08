FROM golang:1.12.6-alpine

MAINTAINER devops787

WORKDIR /go/src/go-rest
COPY . /go/src/go-rest

RUN apk --no-cache add git dep
RUN rm -rf ./vendor && dep ensure && go build -o app

EXPOSE 3000

ENTRYPOINT ["./app"]
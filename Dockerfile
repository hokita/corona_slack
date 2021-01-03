FROM golang:1.15.6-alpine
# For Raspberry Pi
# FROM arm32v6/golang:1.15.6-alpine

ARG SLACK_WEBHOOK_URL

ENV GOBIN /go/bin
ENV GO111MODULE=on
ENV GOPATH=
ENV SLACK_WEBHOOK_URL=$SLACK_WEBHOOK_URL

WORKDIR /go
ADD . /go

RUN go mod tidy
RUN go build -o corona_slack main.go

CMD ./corona_slack

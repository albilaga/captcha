FROM golang:1.17.1-buster

WORKDIR /usr/app

RUN go mod download
RUN go build main.go

CMD [ "./main" ]
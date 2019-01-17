FROM golang:1.11.0-alpine3.8 as app

WORKDIR /go/src/github.com/Bo0km4n/raft-dummy
COPY [".", "."]
RUN go build .
EXPOSE 50051

CMD [""]

FROM golang:1.9.3-alpine3.7

RUN apk update && apk add git

RUN wget https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 -O /usr/bin/dep
RUN chmod a+x /usr/bin/dep

WORKDIR /go/src/github.com/ivanovvpvl/trib
COPY . .

RUN dep ensure

RUN go build

CMD ["./trib"]

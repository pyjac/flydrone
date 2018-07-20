FROM golang:alpine

RUN apk update && \
    apk add git
    
WORKDIR /go/src/flydrone

COPY Gopkg.toml Gopkg.lock ./

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure --vendor-only

COPY . .

CMD ["go", "run", "main.go"]
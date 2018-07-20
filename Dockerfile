FROM golang:alpine

RUN apk update && \
    apk add git
    
WORKDIR /go/src/flydrone

COPY Gopkg.toml Gopkg.lock ./

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure --vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM scratch
COPY --from=0 /go/src/flydrone/main .
ENTRYPOINT ["/main"]


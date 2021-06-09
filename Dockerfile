FROM golang:latest
WORKDIR /go/src/golang_api
ADD . .
ENV CGO_ENABLED=0
RUN go mod vendor && \
    go build

CMD './golang_api'

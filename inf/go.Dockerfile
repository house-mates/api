FROM golang:1.16.3-alpine3.13

RUN go get github.com/gin-gonic/gin
RUN go get github.com/house-mates/api
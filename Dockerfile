FROM golang:1.6
MAINTAINER Nathan Clark <nathan.clark@tokenshift.com>

COPY main.go /go/src/github.com/tokenshift/goprime/
RUN go install github.com/tokenshift/goprime

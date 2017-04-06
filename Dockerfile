FROM golang:1.8

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$PATH
ENV GO15VENDOREXPERIMENT 1
# build directories
# RUN mkdir /app
RUN mkdir /go/src/docker-go-es-nginx-sample

ENV APP_ROOT /go/src/docker-go-es-nginx-sample

WORKDIR $APP_ROOT

COPY . $APP_ROOT

# Go dep!
RUN go get -u github.com/golang/dep/...

RUN dep init -v

RUN dep ensure -update -v

# Build my app
# RUN go build -o /app/main .

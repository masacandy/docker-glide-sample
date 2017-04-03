FROM golang:1.8

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# build directories
# RUN mkdir /app
RUN mkdir /go/src/docker-glide-sample
COPY . /go/src/docker-glide-sample
WORKDIR /go/src/docker-glide-sample

# Go dep!
RUN go get -u github.com/golang/dep/...

# ADD manifest.json manifest.json
# ADD lock.json lock.json

RUN dep ensure -update

# Build my app
# RUN go build -o /app/main .

FROM golang:1.4
CMD ["/go/bin/server"]

COPY . /go/src/github.com/danzel/go-docker-playground

RUN cd /go/src/github.com/danzel/go-docker-playground/server \
        && go get \
        && go install

EXPOSE 35235

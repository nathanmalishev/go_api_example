FROM golang:1.9

ADD . /go/src/github.com/nathanmalishev/go_api_example

RUN go install github.com/nathanmalishev/go_api_example

ENTRYPOINT /go/bin/go_api_example

EXPOSE 8080

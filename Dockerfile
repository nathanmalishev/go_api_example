FROM golang:1.9 as builder
ADD ./ /go/src/github.com/nathanmalishev/go_api_example
WORKDIR /go/src/github.com/nathanmalishev/go_api_example
RUN go get ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


#Container to run
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/nathanmalishev/go_api_example/app .
CMD ["./app"]  


EXPOSE 8080

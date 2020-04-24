FROM golang:1.11.5

# Move current project to a valid go path
COPY . /go/src/github.com/taxky/
WORKDIR /go/src/github.com/taxky/

ENV GOBIN "/go/src/github.com/taxky/build/"
ENV GO111MODULE=on

RUN go mod init
RUN go mod vendor

RUN go install main.go

# Run app in production mode
EXPOSE 7070
CMD ["./build/main"]

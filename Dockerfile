# Specify the base image we need for our go application
FROM golang:1.14
WORKDIR /go/src/app
# Run go build to compile the binary
# executable of our Go program
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Start command which kicks off our
# newly created binary executable
CMD ["app"]
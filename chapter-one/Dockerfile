FROM golang:alpine

RUN mkdir /files

COPY testDocker.go /files
WORKDIR /files
RUN go build -o /files/testDocker testDocker.go
ENTRYPOINT ["/files/testDocker"]
FROM golang:alpine

RUN mkdir /files
COPY . /files
WORKDIR /files

RUN go build -o /files/storage
ENTRYPOINT ["/files/storage"]

FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-s3_client
RUN chmod +x wait-db.sh

RUN go mod download
RUN go build -o pkk ./cmd/pkk/main.go

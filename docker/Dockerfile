FROM golang:1.18.3-alpine

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go build cmd/api/main.go
CMD [ "./main" ]
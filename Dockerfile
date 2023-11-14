FROM golang:1.22.2-alpine
RUN apk update && apk add git
WORKDIR /go/src

CMD ["go", "run", "main.go"]

FROM golang:1.14

WORKDIR /go/src/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

CMD ["go","run","main.go"]
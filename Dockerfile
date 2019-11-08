FROM golang:alpine
WORKDIR /go/src/app
ADD . .

RUN go get && go build

CMD ["./JohnGalt"]


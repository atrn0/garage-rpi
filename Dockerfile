FROM golang:1.18

WORKDIR /go/src/github.com/atrn0/garage-rpi

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app

CMD [ "./app" ]

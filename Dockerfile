FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod download

RUN go build -o /bot

CMD ["/bot"] 
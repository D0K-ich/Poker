FROM golang:alpine as Builder

WORKDIR /build

ADD go.mod .

COPY . .

WORKDIR /build/cmd/app

RUN go mod download

RUN go build -o main main.go

CMD ["./main"]

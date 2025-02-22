FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./main cmd/what_to_watch/main.go

CMD [ "./main"]
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY wait-for-it.sh /wait-for-it.sh

RUN chmod +x /wait-for-it.sh
RUN go build -o main ./cmd/app

EXPOSE 8080

CMD ["/wait-for-it.sh", "db:5432", "--", "./main"]

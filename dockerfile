FROM golang:1.23.9

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN go build -o server main.go

EXPOSE 8010

CMD ["./wait-for-it.sh", "postgres:5432", "--", "./server"]

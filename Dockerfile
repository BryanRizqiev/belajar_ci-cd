FROM golang:alpine

ENV DB_USERNAME="root"
ENV DB_PASSWORD=""
ENV DB_HOST="127.0.0.1"
ENV DB_NAME="data_db"
ENV JWT_SECRET="secret"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

EXPOSE 8080

CMD ["./main"]

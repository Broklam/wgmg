FROM golang:latest

LABEL maintainer="flö <ivfv@yahoo.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

RUN apt-get update && apt-get install -y mariadb-server

ENV MYSQL_ROOT_PASSWORD=1

ENV MYSQL_DATABASE=data

EXPOSE 3184

CMD ["./main"]
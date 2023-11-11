# Start from the latest golang base image
FROM golang:latest

# Add Maintainer info
LABEL maintainer="Broklam <ivnfrv@yahoo.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the Go app
EXPOSE 8080

# Install MariaDB server
RUN apt-get update && apt-get install -y mariadb-server

# Set the MariaDB root password
ENV MYSQL_ROOT_PASSWORD=$DOCKER_DB_PASSWORD

# Set the name of your database
ENV MYSQL_DATABASE=$DOCKER_DB_DATABASE

# Expose port 3184 for MariaDB
EXPOSE 3184

# Run the binary program produced by `go build`
CMD ["./main"]

# Start from the official Golang base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Enable Go modules
ENV GO111MODULE=on

# Initialize Go modules and fetch dependencies
RUN go mod init my-module
RUN go get github.com/InfluxCommunity/influxdb3-go
RUN go get github.com/apache/arrow/go/v12/arrow/flight/flightsql
# RUN go get github.com/apache/arrow/go/arrow/v3
RUN go get github.com/influxdata/line-protocol/v2

# Copy the Go program file into the container
COPY main.go .

# Set ARGs for --build-arg options passed in the build command
ARG INFLUXDB_DATABASE
ARG INFLUXDB_URL
ARG INFLUXDB_TOKEN

# Set run-time ENVs from ARGs
ENV INFLUXDB_DATABASE=${INFLUXDB_DATABASE}
ENV INFLUXDB_URL=${INFLUXDB_URL}
ENV INFLUXDB_TOKEN=${INFLUXDB_TOKEN}

# Compile the Go program
RUN go build -o go-example main.go

# Set the entrypoint to run the compiled program
ENTRYPOINT ["/app/go-example"]
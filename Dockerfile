# Use a Go base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code to the image
COPY . .

# Build the Go application
RUN go build -o main .

# Set the entrypoint to run the application
ENTRYPOINT ["./main"]

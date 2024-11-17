# Use the official Go image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . .

# Set the environment variable for the video directory
ENV VIDEO_DIRECTORY=/app/videos

# Expose the necessary ports (adjust if needed)
EXPOSE 50051 8080

# Build the Go files
RUN go build -o grpc_server server.go && \
    go build -o http_server http_server.go

# Command to run both servers
CMD ["sh", "-c", "echo 'Starting gRPC server...' && ./grpc_server & grpc_pid=$! && sleep 5 && echo 'Starting HTTP server...' && ./http_server & http_pid=$! && wait $grpc_pid && wait $http_pid"]

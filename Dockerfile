# Use the official Golang Alpine image as the base
FROM golang:1.21.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory (where the Dockerfile is located) into the container at /app
COPY . .

# Download dependencies using 'go mod tidy'
RUN go mod tidy

# Build the application, producing an executable named 'main'
RUN go build -o main.

# Inform Docker that the container listens on the specified network ports at runtime. Here we expose port 8080.
EXPOSE 8080

# Define the command to run the application when the container starts
CMD ["./main"]

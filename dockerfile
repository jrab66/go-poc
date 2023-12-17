# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# deleting probably .env on local machines!
RUN rm .env || true
# create dummy .env to remove error on this headless container
RUN touch .env
# Install any needed dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

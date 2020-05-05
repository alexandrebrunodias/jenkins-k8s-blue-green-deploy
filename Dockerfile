# Start from the latest golang base image
FROM golang:1.13.8

# Add Maintainer Info
LABEL maintainer="Alexandre Dias <alexandrebrunodias@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o product main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./product"]
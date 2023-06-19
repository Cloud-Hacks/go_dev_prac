# Start from a base Go image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go program
RUN go build -o app

# Set the command to run the executable
CMD ["./app"]

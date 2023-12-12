# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download dependencies and build the application
RUN go mod download
RUN go build -o main .

# Make port 4444 available to the world outside this container
EXPOSE 4444

# Run the application when the container launches
CMD ["./main"]

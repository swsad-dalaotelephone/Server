# Use an official Golang runtime as a parent image
FROM golang:1.11.5

# Make the working directory
RUN mkdir -p /go/src/github.com/swsad-dalaotelephone/Server

# Set the working directory to /go/src/github.com/swsad-dalaotelephone/Server
WORKDIR /go/src/github.com/swsad-dalaotelephone/Server

# Copy the current directory contents into the container at /go/src/github.com/swsad-dalaotelephone/Server
COPY . /go/src/github.com/swsad-dalaotelephone/Server

# Get all packets and install
RUN go get -u -v 
RUN go install -v

# Make port 8080 available to the world outside this container
EXPOSE 8080

CMD Server run

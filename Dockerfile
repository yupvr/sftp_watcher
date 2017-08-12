# Import base image from https://hub.docker.com/_/golang/
FROM golang
MAINTAINER Yurii Pyvovarov

# Change workdir
WORKDIR /go/src/app
# Copy watcher application into container
COPY watcher.go .

# Install go packages
RUN go-wrapper download
RUN go-wrapper install

# Create directory to mount monitoring directory
RUN mkdir /home/monitor

# Set startup command for  container
CMD ["go-wrapper", "run", "/home/monitor"]

# Pull golang image from docker hub
FROM golang:1.19-alpine

# Install required applications
RUN apk update && \
    apk add make

# Copy the compiled go executable
COPY ./bin/test-xm /

# Set working directory
WORKDIR /

# Copy the source code
COPY . /

ENTRYPOINT ["/test-xm"]

# expose port
EXPOSE 7711

# Pull golang image from docker hub
FROM golang:1.19-alpine

# Install required applications
RUN apk update && \
    apk add make

# Copy the compiled go executable
COPY ./bin/test-xm /
COPY ./xm_app/store/sql/migrations /migrations

# Set working directory
WORKDIR /

ENTRYPOINT ["/test-xm"]

# expose port
EXPOSE 7711

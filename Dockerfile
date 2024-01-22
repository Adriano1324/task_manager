FROM golang:1.21.4

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum /app/
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./ /app/

# Build
RUN go build -o /task_manager

RUN curl -sSf https://atlasgo.sh | sh

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

### STAGE 1: Build ###
# Building the binary of the App
#
#
# We label our stage as 'builder'
FROM golang:1.18 AS builder

ENV PATH=*/go/bin:${PATH}
ENV CGO_ENABLED=0
ENV GO1111MODULE=on

# `boilerplate` should be replaced with your project name
RUN mkdir /go/src/jictyvoo/
WORKDIR /go/src/jictyvoo

# Copy all the Code and stuff to compile everything
COPY . .

WORKDIR /go/src/jictyvoo/server
# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -a -installsuffix cgo -o multiclient_server .
#
#
# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

# General environment variables
ENV APP_KEY=""
ENV APP_URL="localhost:8080"

WORKDIR /app/data

# Create the `public` dir and copy all the assets into it
COPY ./server/config.example.toml ./config.toml

# Copy the generated binary from builder image to execution image
COPY --from=builder /go/src/jictyvoo/server/multiclient_server /bin/multiclient_server

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
ENTRYPOINT ["/bin/multiclient_server", "--host", "0.0.0.0", "--port", "8080"]

CMD ["/bin/multiclient_server"]

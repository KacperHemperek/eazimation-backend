# Specifies a parent image
FROM golang:1.22.0-bullseye

ARG port

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

RUN make build

EXPOSE ${port}

CMD ["./main"]

# Specifies a parent image
FROM golang:1.22.0-bullseye

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

RUN make build

EXPOSE 8080

CMD ["./main"]
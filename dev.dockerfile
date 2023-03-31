# Specifies a parent image
FROM golang:1.20

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

COPY go.mod go.sum /src/

# Installs Go dependencies
RUN go mod download \
  && go mod verify

# Import the code from the context.
COPY . /src

# Builds your app with optional configuration
RUN go build -o /bin/app /src/cmd/app

# Tells Docker which network port your container listens on
EXPOSE 3000

# Specifies the executable command that runs when the container starts
CMD [ "/bin/app" ]
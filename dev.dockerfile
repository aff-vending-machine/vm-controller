# Specifies a parent image
FROM golang:1.20

ENV CGO_ENABLED=1 GO111MODULE=on GOOS=linux GOARCH=arm64

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

COPY go.mod go.sum /src/

# Installs Go dependencies
RUN go mod download
# && go mod verify

# Import the code from the context.
COPY . /src

# Builds your app with optional configuration
RUN go build \
  -mod=mod \
  -a \
  -installsuffix 'static' \
  -o /bin/app \
  /src/cmd/app

# Specifies the executable command that runs when the container starts
CMD [ "/bin/app" ]
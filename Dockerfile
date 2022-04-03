# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.17.8 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Compile the binary
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN go get -v golang.org/x/tools/cmd/goimports
RUN go mod download golang.org/x/net
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o flamingo ./cmd/rest/main.go

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/flamingo /flamingo

# Copy all necessary config files
COPY --from=builder /app/config/config-docker.yaml /config/config.yaml

# If using local database uncomment below
# COPY --from=builder /app/config/config-sample.yaml /config/config-sample.yaml
# RUN sed 's/host: "localhost"/host: "host.docker.internal"/g' /config/config-sample.yaml > /config/config.yaml

# Copy web related files
COPY --from=builder /app/web /web

# Run the web service on container startup.
CMD ["/flamingo"]
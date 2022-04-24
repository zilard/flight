############################
# Multi-stage build - STEP 1 build executable binary
############################
FROM golang:alpine as builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/flight/
COPY *.go .

RUN go env -w GO111MODULE=auto

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/flight


############################
# Multi-stage build - STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/flight /

EXPOSE 8080

# Run the flight binary.
CMD ["/flight"]

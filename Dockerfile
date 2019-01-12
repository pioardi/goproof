# Start from an alpine image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine3.8

# Copy the local package files to the container's workspace.
ADD . /go/src/myapp

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
ENTRYPOINT [ "go" , "run" , "/go/src/myapp/mymain.go" ]

# Document that the service listens on port 8080.
EXPOSE 8080
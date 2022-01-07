#----------------------------------------------------------------
# DEVELOPMENT 
#----------------------------------------------------------------

FROM golang:1.17 AS dev

EXPOSE 8080

WORKDIR /go/src/api

# ----- INSTALL -----

ENV GO111MODULE=on

RUN go mod init

COPY api/go.mod .

# Add missing & remove unused modules

RUN go mod tidy

COPY api/go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY api/ ./

# Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

# Use the excutable
ENTRYPOINT ["/go/bin/air", "-c", ".config/air.toml"]

#----------------------------------------------------------------
# PRODUCTION 
#----------------------------------------------------------------

#
# Builder
#

FROM golang:1.17 AS builder

WORKDIR /go/src/api

# Download necessary Go modules
COPY go.mod .
RUN go mod download

# Copy over the source files
COPY *.go ./

# Build
RUN go build -o /main

#
# Runner
#

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

# Copy from builder the final binary
COPY --from=builder /main /main

USER nonroot:nonroot

ENTRYPOINT ["/main"]
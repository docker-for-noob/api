#----------------------------------------------------------------
# DEVELOPMENT 
#----------------------------------------------------------------

FROM golang:1.17 AS dev

EXPOSE 8080

WORKDIR /go/src/api

# ----- INSTALL -----

ENV GO111MODULE=on

RUN go mod init

COPY go.mod .

# Add missing & remove unused modules

RUN go mod tidy

COPY go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY . ./

# Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

# Use the excutable
ENTRYPOINT ["/go/bin/air", "-c", ".config/air.toml"]
#----------------------------------------------------------------
# DEVELOPMENT 
#----------------------------------------------------------------

FROM golang:1.17

EXPOSE 8080

WORKDIR /api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY . ./

## Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air


## Use the excutable
ENTRYPOINT ["/go/bin/air", "-c", ".config/air.toml"]
FROM golang:1.17

EXPOSE 8080

WORKDIR /go/src/api

ENV GO111MODULE=on


RUN go mod init api

COPY go.mod ./

RUN go mod tidy

COPY go.sum ./
RUN go mod download -x


COPY . ./

## Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

RUN go mod vendor


## Use the executable
ENTRYPOINT ["/go/bin/air", "-c", ".config/air.toml"]
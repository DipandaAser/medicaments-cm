FROM golang:latest AS build

LABEL authors="Aser Dipanda<aserdipanda@gmail.com>"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

#fetch dependencies
COPY go.* ./
RUN go mod download

#build the binary
COPY . .
RUN go build -o bin/app-bin .

FROM alpine

WORKDIR /app
COPY --from=build /build/bin/app-bin /app/

EXPOSE 7000

ENTRYPOINT ["/app/app-bin"]
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add build-base
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app

ENTRYPOINT ./app
LABEL Name=veebiprogrammeerimine Version=0.0.1
EXPOSE 8080

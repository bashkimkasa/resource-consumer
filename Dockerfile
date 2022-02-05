
###############
# build stage #
###############
FROM golang:alpine AS builder
WORKDIR /go/src/app

ENV GO111MODULE=on
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./consumer

###############
# final stage #
###############
FROM alpine:latest

RUN apk --no-cache add ca-certificates stress-ng

COPY --from=builder /go/src/app/consumer /app/consumer
WORKDIR /app
LABEL Name=resource-consumer
EXPOSE 8080
ENTRYPOINT ["./consumer"]
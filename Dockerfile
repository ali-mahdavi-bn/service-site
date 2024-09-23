#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=cosmicgo Version=0.0.1
EXPOSE 3000

FROM builder AS dev
RUN go install github.com/mitranim/gow@latest
CMD ["gow", "run", "."]

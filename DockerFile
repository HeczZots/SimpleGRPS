# doesnt work
FROM golang:1.21 AS client-builder
WORKDIR /app
COPY cmd/client .
RUN go build -o client

FROM golang:1.21 AS server-builder
WORKDIR /app
COPY cmd/server .
RUN go build -o server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=client-builder /app/client .
COPY --from=server-builder /app/server .
CMD ["./server","./client"]
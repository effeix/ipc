FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ipc .

FROM alpine:latest

COPY --from=builder /app/ipc /usr/local/bin/ipc

RUN apk add --no-cache bash && \
    rm -rf /app

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]

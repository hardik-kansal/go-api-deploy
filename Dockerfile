FROM golang:alpine3.20 AS builder
RUN apk add --no-cache curl tar
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
RUN chmod +x start.sh wait-for.sh
EXPOSE 8080
ENTRYPOINT ["/wait-for.sh"]
CMD ["./main"]
FROM golang:alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
# ./main requires env variables
COPY app.env .
EXPOSE 8080

CMD ["./main"]
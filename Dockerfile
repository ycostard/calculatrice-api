FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM scratch
COPY --from=builder /app/main /main

EXPOSE 8080
ENTRYPOINT ["/main"]

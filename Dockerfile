FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o beaver main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/beaver ./beaver
EXPOSE 9001
USER nobody
ENTRYPOINT ["./beaver"]

FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux \ 
    go build -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o gogomanager ./cmd/gogomanager/main.go

FROM alpine:3.19
RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D appuser
WORKDIR /app

COPY --from=builder /app/gogomanager .
COPY .env .
COPY internal/database/schema ./internal/database/schema

RUN chown -R appuser:appuser /app
USER appuser

# match the number of cores and memory limit with the machine specs
ENV GOMAXPROCS=1 
ENV GOGC=100
ENV GOMEMLIMIT=750MiB

EXPOSE 8080

CMD ["./gogomanager"]
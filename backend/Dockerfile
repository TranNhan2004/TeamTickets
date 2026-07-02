# ======================
# Stage 1: Build Go app
# ======================
FROM golang:1.23-alpine AS build

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ticket-api ./cmd/api


# ======================
# Stage 2: Run Go binary
# ======================
FROM alpine:3.20 AS run

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=build /app/ticket-api ./ticket-api

EXPOSE 8080

CMD ["./ticket-api"]
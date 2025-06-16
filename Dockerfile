FROM golang:1.24.3-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o online-restaurant ./backend

FROM scratch
COPY --from=builder /src/online-restaurant /online-restaurant
EXPOSE 8080
ENTRYPOINT ["/online-restaurant"]
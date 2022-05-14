FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.* .
RUN go mod tidy

COPY . .
RUN go build -o server ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]
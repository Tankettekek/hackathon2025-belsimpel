FROM golang:latest

# Stage 1 : Build the Go application
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# Stage 2 : Create a small image to run the application
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=0 /app/main .
EXPOSE 8080
CMD ["./main"]
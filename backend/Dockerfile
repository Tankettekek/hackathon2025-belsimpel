FROM golang:latest

# Stage 1 : Build the Go application
WORKDIR /app
COPY . .
COPY src/templates /app/src/templates
RUN go mod tidy
RUN go build -o main .

# Stage 2 : Create a small image to run the application
FROM golang:latest
WORKDIR /app
COPY --from=0 /app/main .
COPY --from=0 /app/src/templates /app/src/templates
EXPOSE 8080
CMD ["./main"]
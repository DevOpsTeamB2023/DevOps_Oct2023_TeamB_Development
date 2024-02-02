# Build Stage
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go build -o myapp

# Final Stage
FROM heroku/go:latest
WORKDIR /app
COPY --from=build /app/myapp .
CMD ["./myapp"]

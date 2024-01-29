# Stage 1: Build Go application
FROM golang:1.21.3 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/console-app ./console

# Stage 2: Build final image with Alpine and your Go application
FROM alpine:3.15 as final
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app
COPY --from=builder /app/console-app /app/console-app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER app
ENTRYPOINT ["/app/console-app"]

# Stage 3: Build final image with MySQL
FROM mysql:latest
ENV MYSQL_ROOT_PASSWORD=root_password
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=user_password
ENV MYSQL_DATABASE=mydatabase

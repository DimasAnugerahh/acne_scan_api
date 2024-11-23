# FROM golang:1.23

# WORKDIR /usr/src/app

# # pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# COPY . .
# RUN go mod tidy
# RUN go build -v -o /usr/local/bin/app ./cmd/main.go

# # Install migrate tool
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
#   && tar -xvzf migrate.tar.gz \
#   && mv migrate /usr/local/bin/migrate \
#   && chmod +x /usr/local/bin/migrate


# # Copy the migrations into the Docker image
# COPY internal/infrastructure/database /app/internal/infrastructure/database


# CMD ["app"]


# Use a specific Go version
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the project files
COPY . .

# Install migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
  && tar -xvzf migrate.tar.gz \
  && mv migrate /usr/local/bin/migrate \
  && chmod +x /usr/local/bin/migrate

# Ensure dependencies are installed
# RUN sudo go mod tidy

# Build the application
RUN sudo go build -v -o /usr/local/bin/app cmd/main.go

# Run the application
CMD ["/usr/local/bin/app"]


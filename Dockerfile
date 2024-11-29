FROM golang:1.23

WORKDIR /usr/src/app

RUN go clean -modcache

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go mod tidy

COPY . .

# COPY ./acne-scan-bucket.json /usr/src/app/


RUN go build -v -o /usr/local/bin/app ./cmd/main.go

# Install migrate tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
  && tar -xvzf migrate.tar.gz \
  && mv migrate /usr/local/bin/migrate \
  && chmod +x /usr/local/bin/migrate


# Copy the migrations into the Docker image
COPY internal/infrastructure/database /app/internal/infrastructure/database

ENV DB_USERNAME=root
ENV DB_PASSWORD=acne-scan-db-password
ENV DB_NAME=acne-scan
ENV DB_PORT=3306
ENV DB_HOST=/cloudsql//acnescan-final:asia-southeast2:acne-scan-sql

CMD ["app"]

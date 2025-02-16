# Stage 1: Build
FROM golang:1.23.2 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main

# Stage 2: Final image
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build /app/main .

# Expose port 3000 (optional, helps clarity but not mandatory)
EXPOSE 3000

# Just start our binary; Railway sets env variables at runtime
ENTRYPOINT ["/app/main"]
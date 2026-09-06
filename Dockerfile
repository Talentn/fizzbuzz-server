FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /fizzbuzz-server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /fizzbuzz-server /fizzbuzz-server
EXPOSE 8080
ENTRYPOINT ["/fizzbuzz-server"]
FROM golang:1.22-alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/bin/golang-echo-api

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /app/bin/. .
EXPOSE 8000
CMD ["./golang-echo-api"]

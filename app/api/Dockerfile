FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest as prod
EXPOSE 8080
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder app ./
CMD ["./app"]

FROM golang:1.13

WORKDIR /app
COPY main.go ./
COPY go.mod ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/go .
RUN chmod +x /app/go
EXPOSE 6001/tcp
CMD ["./go"]
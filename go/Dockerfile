FROM golang:1.13

WORKDIR /app
COPY main.go ./
COPY go.mod ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build
RUN chmod +x /app/go

EXPOSE 6001/tcp 
ENTRYPOINT ["/app/go"]

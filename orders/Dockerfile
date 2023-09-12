FROM golang:1.21.0
WORKDIR /app
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/orders ./cmd/orders/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/build/orders ./
CMD ["/root/orders"]

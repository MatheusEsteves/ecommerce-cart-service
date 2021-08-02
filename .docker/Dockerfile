FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
ADD . .

RUN go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o /tmp/ecommerce-cart-service

FROM scratch

WORKDIR /app
COPY --from=builder /tmp/ecommerce-cart-service /app/ecommerce-cart-service

ENTRYPOINT ["/app/ecommerce-cart-service"]
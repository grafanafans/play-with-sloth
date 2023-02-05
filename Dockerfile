FROM golang:alpine3.16 AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN GOPROXY=off go build -o myservice -mod=vendor cmd/main.go

FROM alpine
WORKDIR /app 
COPY --from=builder /build/myservice /app/myservice
ENTRYPOINT ["./myservice"]
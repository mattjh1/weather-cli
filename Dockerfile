FROM golang:1.21-alpine AS build

WORKDIR /go/src/app

COPY . .

RUN go build -o weather.go .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/app/weather .

CMD ["./weather"]

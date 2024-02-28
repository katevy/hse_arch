FROM golang:1.21-alpine AS build

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app

WORKDIR /app
COPY . .

RUN go mod download
RUN go test ./...
RUN go build -o ./api main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=build /app/api .
COPY --from=build /app/database.db .

EXPOSE 8080

CMD ["./api"]


FROM golang:1.13-alpine AS build-env
WORKDIR /app
COPY . .
RUN go build -o build/service cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /app/internal/ /app/internal/
COPY --from=build-env /app/build/service /app/service
RUN chmod +x /app/service
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 8000
CMD ["./service"]

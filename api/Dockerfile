FROM golang:latest as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY . /app
COPY --from=builder /build/main /app
WORKDIR /app
EXPOSE 3000
CMD ["ls"]
CMD ["./main"]

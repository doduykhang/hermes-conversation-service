FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o app ./cmd/main

RUN chmod +x /app/app

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/app /app
COPY --from=builder /app/db /db
COPY --from=builder /app/config.json /config.json

CMD [ "/app/app" ]

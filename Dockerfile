FROM golang:1.16-alpine as builder

WORKDIR /root

COPY . .

RUN go mod download

RUN go build -o /app

FROM alpine

RUN apk add --no-cache tzdata

WORKDIR /root

COPY --from=builder /app .

CMD [ "/root/app" ]


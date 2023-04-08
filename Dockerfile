FROM golang:1.20-alpine3.17 AS builder

COPY . /MyFit/
WORKDIR /MyFit/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /MyFit/bin/bot .
COPY --from=0 /MyFit/configs configs/
COPY --from=0 /MyFit/pics pics/

EXPOSE 80

CMD ["./bot"]
FROM golang:1.17.8-alpine3.14 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/bowl

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/bowl .

FROM alpine:3.15

RUN apk add ca-certificates

COPY --from=build_base /tmp/bowl/out/bowl /app/bowl

WORKDIR /app

EXPOSE 8080

CMD ["./bowl"]

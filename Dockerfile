FROM golang:1.17 AS builder

RUN go version

COPY . /github.com/salesforceanton/products-data-manager/
WORKDIR /github.com/salesforceanton/products-data-manager/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/salesforceanton/products-data-manager/.bin/app .

CMD ["./app"]
FROM golang:1.22.4-alpine

RUN apk update && apk add --no-cache \
  python3 \
  py3-pip

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/api

CMD ["/app/app"]

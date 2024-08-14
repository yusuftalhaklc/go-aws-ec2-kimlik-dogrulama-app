FROM golang:1.21.0

WORKDIR /app

COPY . .

RUN go build -o kimlik-dogrulama api

ENV GIN_MODE=release

CMD ["./kimlik-dogrulama"]

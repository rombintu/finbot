FROM golang:1.18-alpine
COPY . /opt/finbot/
WORKDIR /opt/finbot
RUN go build -o finbot cmd/main.go
CMD './finbot'
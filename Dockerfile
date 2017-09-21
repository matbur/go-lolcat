FROM golang:alpine

WORKDIR /app
COPY . /app
RUN go build -o lolcat
CMD cat main.go | ./lolcat

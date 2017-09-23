FROM golang

WORKDIR /app
RUN go get github.com/jessevdk/go-flags
CMD go build -o lolcat ./src

# go-lolcat

LolCat written in Golang based on Python version linked below.

#### Other implementations:
- Ruby: https://github.com/busyloop/lolcat
- Python: https://github.com/tehmaze/lolcat

## Usage locally:

### Installation:
```
go build -o lolcat
```

### Running:
```
cat main.go | ./lolcat
```

## Usage in Docker:

### Building an image:
```
docker build -t lolcat .
```

### Running container:
```
docker run --rm lolcat
```

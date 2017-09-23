# go-lolcat

LolCat written in Golang based on Python version linked below.

#### Other implementations:
- Ruby: https://github.com/busyloop/lolcat
- Python: https://github.com/tehmaze/lolcat

## Usage locally:

### Installation:
```
go build -o lolcat ./src
```

### Running:
```
./lolcat -h | ./lolcat
```

## Usage in Docker:

### Building an image:
```
make build-image
```

### Running container:
```
make
```

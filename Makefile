install:
	go install

clean:
	rm ./bin/*

linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/zeroconf-beacon-linux-amd64

osx:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/zeroconf-beacon-darwin-amd64

all: install linux osx
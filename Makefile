build:
	go build -o cmd ./app

run: clean build
	./cmd

test:
	go test ./...

clean:
	rm -f cmd
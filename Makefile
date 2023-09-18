build:
	go build -o ./app/cmd ./app

run: clean build
	./app/cmd

test:
	go test ./...

test_cover:
	go test -cover ./...

clean:
	rm -f ./app/cmd
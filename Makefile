build:
	go build -o bin/fuzzygit cmd/fuzzygit/main.go

run:
	go run cmd/fuzzygit/main.go

build-test:
	go build -o /tmp/testdir/fuzzygit cmd/fuzzygit/main.go
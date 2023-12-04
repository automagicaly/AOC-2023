default: bin compile

bin:
	mkdir bin

compile:
	go build -o bin ./cmd/...

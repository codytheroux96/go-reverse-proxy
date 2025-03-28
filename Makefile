# Format the code using gofmt
clean:
	gofmt -w .

# Build the project
build:
	go build -o go-reverse-proxy ./cmd/go_reverse_proxy

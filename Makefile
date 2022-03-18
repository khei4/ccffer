run: 
	go run main.go ./test
test:
	go test ./... -v -race
	
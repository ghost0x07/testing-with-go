test-all:
	go test -cover -v ./...
test-structs:
	go test -cover -v ./structs/
test-maps:
	go test -cover -v ./maps/
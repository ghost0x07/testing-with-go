test-all:
	go test -cover -v ./...
test-structs:
	go test -cover -v ./structs/
test-maps:
	go test -cover -v ./maps/
test-di:
	go test -cover -v ./di/
test-mocks:
	go test -cover -v ./mocks/
test-all:
	go test  -cover -v -bench="." ./...
test-structs:
	go test -cover -v ./structs/
test-maps:
	go test -cover -v ./maps/
test-di:
	go test -cover -v ./di/
test-mocks:
	go test -cover -v ./mocks/
test-concurrency:
	go test -race -cover -v -bench=. ./concurrency/
test-select:
	go test -cover -v ./selector/
test-reflection:
	go test -cover -v ./reflection/
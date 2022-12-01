.PHONY: test day1

test:
	go test -coverprofile=cover.out ./...

day1:
	go run cmd/day1/*.go
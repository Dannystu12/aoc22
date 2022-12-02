.PHONY: test day1 day2

test:
	go test -coverprofile=cover.out ./...

day1:
	go run cmd/day1/*.go

day2:
	go run cmd/day2/*.go
.PHONY: test day01 day02 day03 day04 day05 day06 day07 day08 day09 day10

test:
	go test -coverprofile=cover.out -covermode=atomic ./...

day01:
	go run cmd/day01/*.go

day02:
	go run cmd/day02/*.go

day03:
	go run cmd/day03/*.go

day04:
	go run cmd/day04/*.go

day05:
	go run cmd/day05/*.go

day06:
	go run cmd/day06/*.go

day07:
	go run cmd/day07/*.go

day08:
	go run cmd/day08/*.go

day09:
	go run cmd/day09/*.go

day10:
	go run cmd/day10/*.go
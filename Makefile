.PHONY: test day1 day2 day3 day4 day5 day6

test:
	go test -coverprofile=cover.out ./...

day1:
	go run cmd/day1/*.go

day2:
	go run cmd/day2/*.go

day3:
	go run cmd/day3/*.go

day4:
	go run cmd/day4/*.go

day5:
	go run cmd/day5/*.go

day6:
	go run cmd/day6/*.go
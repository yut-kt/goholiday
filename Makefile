updateHolidays:
	bash script/holiday_downloader.sh
	go run gen/jp_generator.go

test:
	go test -v ./...

bench:
	go test -bench . -benchmem -count 5 -run none

cov:
	go test -cover

fmt:
	golangci-lint run ./...
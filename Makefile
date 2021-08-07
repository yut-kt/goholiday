updateHolidays:
	bash script/holiday_downloader.sh
	go run gen/jp_generator.go

test:
	go test -v ./...

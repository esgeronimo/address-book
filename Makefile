default: build

init:
	@rm -rf build
	@mkdir build

build: init test
	go build

run: build test
	go run main.go

test: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out ./...
	
coverage: init test
	go tool cover -html=build/cover.out 

test_integ:
	newman run esgeronimo-address-book.postman_collection.json

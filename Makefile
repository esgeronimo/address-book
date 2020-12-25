default: build

init:
	@rm -rf build
	@mkdir build

build: init test
	go build

test: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out ./...
	
coverage: init test
	go tool cover -html=build/cover.out 
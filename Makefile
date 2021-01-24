VERSION := develop
IMAGE := esgeronimo/address-book:${VERSION}

default: build

init:
	@rm -rf build
	@mkdir build

build: init test
	docker build -t ${IMAGE} .

deploy: build
	kubectl create deployment address-book --image=${IMAGE}

test: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out ./...
	
coverage: init test
	go tool cover -html=build/cover.out 

test_integ:
	newman run esgeronimo-address-book.postman_collection.json

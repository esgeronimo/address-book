VERSION := develop
IMAGE := esgeronimo/address-book:${VERSION}

default: build

init:
	@rm -rf build
	@mkdir build

build: init test
	docker build -t ${IMAGE} .

deploy: build
	sed 's/{{ .Values.image.tag }}/develop/g' k8/address-book-deployment.yaml | kubectl apply -f -

test: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out ./...
	
coverage: init test
	go tool cover -html=build/cover.out 

test_integ:
	newman run postman/esgeronimo-address-book.postman_collection.json

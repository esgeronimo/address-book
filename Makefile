VERSION := develop
K8_SVC_NAME := address-book
IMAGE := esgeronimo/address-book:${VERSION}

default: build

init:
	@rm -rf build
	@mkdir build

build: init test
	docker build -t ${IMAGE} .
	sed 's/{{ .Values.image.tag }}/develop/g;s/{{ .Values.serviceName }}/${K8_SVC_NAME}/g' k8/address-book-deployment.yaml > build/address-book-deployment.yaml

deploy: build
	kubectl apply -f build/address-book-deployment.yaml
	kubectl apply -f k8/address-book-configmap.yaml

test: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out ./...
	
coverage: init test
	go tool cover -html=build/cover.out 

test_integ:
	# Change SED to create file under build/ directory
	$(eval NODE_PORT = `kubectl get svc -o=jsonpath='{.items[?(@.metadata.name=="${K8_SVC_NAME}")].spec.ports[?(@.port==8080)].nodePort}'`)
	sed "s/{{port}}/$(NODE_PORT)/g" postman/local.postman_environment.json > build/local.postman_environment.json
	newman run postman/esgeronimo-address-book.postman_collection.json -e build/local.postman_environment.json

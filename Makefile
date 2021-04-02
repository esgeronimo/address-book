VERSION := develop
K8_SVC_NAME := address-book
IMAGE := esgeronimo/address-book:${VERSION}
DB_IMAGE := esgeronimo/address-book-mysql:${VERSION}

default: build

init:
	@rm -rf build
	@mkdir build

build: init build_db build_app

build_app: init test_app
	docker build -t ${IMAGE} .
	sed 's|{{ .Values.image }}|${IMAGE}|g;s|{{ .Values.serviceName }}|${K8_SVC_NAME}|g' k8/address-book-deployment.yaml > build/address-book-deployment.yaml

test_app: init
	go test -v -coverpkg=./... -coverprofile=build/cover.out -json ./... > build/test-report.json

coverage: init test_app
	go tool cover -html=build/cover.out	

build_db: init
	docker build -t ${DB_IMAGE} mysql
	sed 's|{{ .Values.image }}|${DB_IMAGE}|g' k8/mysql-deployment.yaml > build/mysql-deployment.yaml

deploy: build
	kubectl apply -f build/mysql-deployment.yaml
	kubectl apply -f build/address-book-deployment.yaml
	kubectl apply -f k8/address-book-configmap.yaml

test_integ:
	$(eval NODE_PORT = `kubectl get svc -o=jsonpath='{.items[?(@.metadata.name=="${K8_SVC_NAME}")].spec.ports[?(@.port==8080)].nodePort}'`)
	sed "s/{{port}}/$(NODE_PORT)/g" postman/local.postman_environment.json > build/local.postman_environment.json
	newman run postman/esgeronimo-address-book.postman_collection.json -e build/local.postman_environment.json

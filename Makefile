GOBIN=$(shell pwd)/bin

generate: gen-endpoints gen-services

gen-endpoints:
	go generate ./models/endpoints

gen-services:
	go generate ./service

install-deps:
	mkdir -p ${GOBIN}
	GOBIN=${GOBIN} go install --tags codegen github.com/aws/aws-sdk-go-v2/private/model/cli/gen-endpoints
	GOBIN=${GOBIN} go install --tags codegen github.com/aws/aws-sdk-go-v2/private/model/cli/gen-api
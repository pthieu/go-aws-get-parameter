exec_name=ssm_get_parameter

build-linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $(exec_name) ssm_get_parameter.go

build-mac:
	go build -o $(exec_name) ssm_get_parameter.go

develop: build-docker
	docker run -it --rm \
	-v $(shell pwd):$(shell pwd) \
	-e GOOS=linux \
	-e GOARCH=386 \
	-e CGO_ENABLED=0 \
	-e AWS_REGION=us-west-2 \
	-e IAM_ROLE \
	ssm-get-parameter sh

build-docker:
	docker build -t ssm-get-parameter .

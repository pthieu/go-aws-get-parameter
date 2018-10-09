exec_name=ssm_get_parameter

build-linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $(exec_name) ssm_get_parameter.go

build-mac:
	go build -o $(exec_name) ssm_get_parameter.go

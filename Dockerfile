FROM golang:alpine as BUILDER
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR /app
ENV GO111MODULE=on
RUN go mod download
COPY . /app

FROM BUILDER as BINARY
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ssm_get_parameter

FROM alpine
ENV AWS_REGION=us-west-2
COPY --from=BINARY /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BINARY /etc/passwd /etc/passwd
COPY --from=BINARY /app/ssm_get_parameter /ssm_get_parameter
ENTRYPOINT ["/ssm_get_parameter"]

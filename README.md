## AWS Get Parameter

### TLDR
Download the executable [here](wget https://github.com/pthieu/go-aws-get-parameter/releases/download/v1.0.0-arm64/ssm_get_parameter) (or whatever release you want)

### Build
Change `GOOS` and `GOARCH` to your desired platform.

**x86_64**
```
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ssm_get_parameter
```

**arm64**
```
GOOS=linux GOARCH=arm64 go build -o ssm_get_parameter
```


### Docker
**Dockerfile**
```docker

ARG ARG_SSM_BASE_PATH

FROM node:20-alpine AS base
WORKDIR /app

# Set env for image
ARG ARG_SSM_BASE_PATH
ENV SSM_BASE_PATH=$ARG_SSM_BASE_PATH
RUN echo "SSM_BASE_PATH is set to: $SSM_BASE_PATH"

# Download binary
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN wget https://github.com/pthieu/go-aws-get-parameter/releases/download/v1.0.0-arm64/ssm_get_parameter
RUN ["chmod", "+x", "./ssm_get_parameter"]

COPY --from=builder /app/docker-entrypoint.sh ./

ENTRYPOINT [ "sh", "docker-entrypoint.sh" ]
CMD ["node", "server.js"]
```

**docker-entrypoint.sh**
```sh
#!/bin/bash
set -e

ssm_available() {
  if [ -z "$SSM_BASE_PATH" ]; then
    return 1
  fi

  return 0
}

# Read: https://github.com/pthieu/go-aws-get-parameter
export_ssm_params_go() {
  eval $(./ssm_get_parameter --path ${SSM_BASE_PATH})
  exec "$@"
}

main() {
  if ssm_available; then
    echo "Info: Loading SSM Parameters" >&2
    export_ssm_params_go "$@"
  fi

  echo "Info: Starting ..." >&2
  exec "$@"
}

main "$@"
```

#### Get a Single Param

```sh
export MY_SECRET=$(./ssm_get_parameter --name /secret/path/var)
echo $MY_SECRET
/# secret
```

#### Get all params in a path and set as ENV vars

```sh
# Set all vars in a path
eval $(./ssm_get_parameter --path /secret/path)
echo $SECRET_1
/# secret1
echo $SECRET_2
/# secret2
```

### Flags
| Flag Name | Default | Description |
| --------- | ------- | ----------- |
| `--region` | `us-west-2` | AWS Region to get Param from |
| `--name`| `None` | The SSM parameter name |
| `--path` | `None` | The SSM parameter path |

### Why?

Sometimes you inherit legacy software or have to shim in secrets because people don't think about storing them securely when making apps. This binary allows you to add that functionality in a shell script so that you can enable this functionality in your deployable.

### AWS Get Parameter

##### TLDR
Download the executable [here](https://github.com/pthieu/go-aws-get-parameter/blob/master/ssm_get_parameter).

#### Build
```
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ssm_get_parameter
```
Change `GOOS` and `GOARCH` to your desired platform.

##### Get a Single Param

```sh
wget https://github.com/pthieu/go-aws-get-parameter/blob/master/ssm_get_parameter
chmod +x ssm_get_parameter

# If running alpine linux, run `apk add ca-certificates`
export MY_SECRET=$(./ssm_get_parameter --name /secret/path/var)

echo $MY_SECRET
# SECRET
```

##### Get all params in a path and set as ENV vars

```sh
# Get the executable
wget https://github.com/pthieu/go-aws-get-parameter/blob/master/ssm_get_parameter
chmod +x ssm_get_parameter

# If running alpine linux, run `apk add ca-certificates`

# set all vars in a path
eval $(./ssm_get_parameter --path /secret/path)
echo $SECRET_PATH_RESULT_1
/# secret1
echo $SECRET_PATH_RESULT_2
/# secret2

```

#### Flags
| Flag Name | Default | Description |
| --------- | ------- | ----------- |
| `--region` | `us-west-2` | AWS Region to get Param from |
| `--name`| `None` | The SSM parameter name |
| `--path` | `None` | The SSM parameter path |

#### Why?

Sometimes you inherit legacy software or have to shim in secrets because people don't think about storing them securely when making apps. This binary allows you to add that functionality in a shell script so that you can enable this functionality in your deployable.

### AWS Get Parameter

##### Tl;DR;
Download the executable [here](https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter).

##### Get a Single Param

```sh
wget https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter
chmod +x ssm_get_parameter

# If running alpine linux, run `apk add ca-certificates`
export MY_SECRET=$(./ssm_get_parameter --name /dev/my-secret)

echo $MY_SECRET
# SECRET
```

##### Get all params in a path and set as ENV vars

```sh
# Get the executable
wget https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter
chmod +x ssm_get_parameter

# If running alpine linux, run `apk add ca-certificates`

# set all vars in a path
eval $(./ssm_get_parameter --path /<secret-path>)
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

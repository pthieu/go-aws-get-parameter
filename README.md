### AWS Get Parameter

##### Tl;DR;
Download the executable [here](https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter).

##### Get a Single Param

```sh
wget https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter
chmod +x ssm_get_parameter
export MY_SECRET=$(./ssm_get_parameter --name /dev/my-secret)

echo $MY_SECRET
# SECRET
```

##### Get all params in a path

```sh
# Get the executable
wget https://s3-us-west-2.amazonaws.com/kloudcover-tools/binaries/ssm_get_parameter
chmod +x ssm_get_parameter

# set all vars in a path
eval $(./ssm_get_paramter --path /<secret-path>)
echo $SECRET_PATH_RESULT_1
/# secret1
echo $SECRET_PATH_RESULT_2
/# secret2

```


This compiles code so that you can use a binary to get AWS SSM secrets without having to install anything else. If you want to build, make sure you have docker installed.

### Usage

Make sure to set `AWS_REGION` to your proper region in the shell.

##### Single Value
```
./ssm_get_parameter --name <ssm name>
```

##### Set Many Environment Variables
```
./ssm_get_parameter --path /mypath
```

##### Concrete example
```sh
export BOSTON=$(./ssm_get_parameter --name jira.url)`
echo $BOSTON
/# https://atlassian.boston.com
```

Or use the multi env var feature by path
```sh
eval $(./ssm_get_paramter --path /<secret-path>)
echo $SECRET_PATH_RESULT_1
/# secret1
echo $SECRET_PATH_RESULT_2
/# secret2
```
### To build

Linux
```
make build-linux
```

Mac
```
make build-mac
```

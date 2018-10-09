# AWS SSM get secret

This is a go binary to get a secret and populate an environment variable
with this said secret.

Make sure to set `AWS_REGION` because otherwise it will fail.

Usage:
`./ssm_get_parameter --name <secret-key>`
`./ssm_get_parameter --path /mypath`

Concrete example
`export BOSTON=$(./ssm_get_parameter --name jira.url)`

`echo $BOSTON`
`https://atlassian.spscommerce.com`

Or

`eval $(./ssm_get_paramter --path /<secret-path>)`
`echo $SECRET_PATH_RESULT`

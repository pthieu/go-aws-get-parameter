# AWS SSM get secret

This is a go binary to get a secret and populate an environment variable
with this said secret.

Make sure to set `AWS_REGION` because otherwise it will fail.

Usage:
`./ssm_get_parameter <secret-key>`


Concrete example
`export BOSTON=$(./ssm_get_parameter jira.url)`

`echo $BOSTON`
`https://atlassian.spscommerce.com`

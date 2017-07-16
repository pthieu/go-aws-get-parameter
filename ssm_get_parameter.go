package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {

	if len(os.Args) != 2 {
		exitErrorf("SSM Key name required\nUsage: %s key_name",
			os.Args[0])
	}

	key_name := os.Args[1]

	// Create S3 service client
	svc := ssm.New(session.New())

	params := &ssm.GetParameterInput{
		Name:           aws.String(key_name),
		WithDecryption: aws.Bool(true),
	}

	// Get the list of items
	resp, err := svc.GetParameter(params)

	if err != nil {
		exitErrorf("Unable to get key %q, %v", key_name, err)
	}

	fmt.Println(*resp.Parameter.Value)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

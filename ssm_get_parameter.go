package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	var region = flag.String("region", "us-west-2", "AWS Region")
	var path = flag.String("path", "", "parameter path")
	var param_name = flag.String("name", "", "single parameter value")
	flag.Parse()

	// Make sure users set values
	if *path == "" && *param_name == "" {
		exitErrorf("Make sure you set --path or --name as arguments")
	}
	// Create SSM service client
	svc := ssm.New(session.New(&aws.Config{Region: aws.String(*region)}))

	if *path == "" {
		params := &ssm.GetParameterInput{
			Name:           aws.String(*param_name),
			WithDecryption: aws.Bool(true),
		}
		resp, err := svc.GetParameter(params)
		if err != nil {
			exitErrorf("Unable to get key %q, %v", *param_name, err)
		}
		fmt.Println(*resp.Parameter.Value)

	} else {
		params := &ssm.GetParametersByPathInput{
			Path:           aws.String(*path),
			Recursive:      aws.Bool(true),
			WithDecryption: aws.Bool(true),
		}
		resp, err := svc.GetParametersByPath(params)
		if err != nil {
			exitErrorf("Unable to get key %q, %v", *path, err)
		}

		for _, v := range resp.Parameters {
			the_key := strings.Split(*v.Name, "/")
			fmt.Printf("export %s=%s\n", the_key[len(the_key)-1], *v.Value)
		}
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

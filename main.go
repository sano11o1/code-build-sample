package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() (string, error) {
	return fmt.Sprintf("build by code build🏢🏢🚀"), nil
}

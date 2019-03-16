package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	fmt.Println("SECRETS MANAGER THIS IS GETTING RUN")
}

func main() {
	lambda.Start(handler)
}

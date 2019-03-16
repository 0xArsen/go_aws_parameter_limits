package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	log.Println("START OF FUNCTION")

	setEnv() //Setting pameter environment variable, for further usage
	log.Println(os.Getenv("PARAM_NAME"))

	log.Println("END OF FUNCTION")
}

func setEnv() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		log.Fatal(err)
	}

	ssmvc := ssm.New(sess)
	keyname := os.Getenv("YOUR_PARAMETER_PATH")
	paramInput := &ssm.GetParameterInput{
		Name:           aws.String(keyname),
		WithDecryption: aws.Bool(true)}

	param, err := ssmvc.GetParameter(paramInput)
	if err != nil {
		log.Printf("Param store response: %s, name(parameter): %s\n", err.Error(), keyname)
	}
	os.Setenv("PARAM_NAME", *param.Parameter.Value)

}

func main() {
	lambda.Start(handler)
}

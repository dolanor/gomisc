package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	/*
		This is an example showing how to use the aws golang sdk
		with digitalocean spaces. To use replace the variables
		`endpoint` and `region` with your digitalocean endpoint
		and `myBucket` with your D.O. bucket name.

		The program when run with `go run main.go` will upload `main.go`
		to your digital ocean space `myBucket`.

		To install aws sdk `go get -u github.com/aws/aws-sdk-go`
		Note: your API key for D.O. must be in ~/.aws/credentials
	*/
	err := Uploader()
	fmt.Println(err)
}

func Uploader() error {
	// https://jto.nyc3.digitaloceanspaces.com
	// The session the S3 Uploader will use
	endpoint := "ams3.digitaloceanspaces.com"
	region := "ams3"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}))

	svc := s3.New(sess)
	input := &s3.CreateBucketInput{
		Bucket: aws.String("uniqueid-somewhat-to-test-whatsup3"),
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	fmt.Println(result)
	return nil
}

# Links to use as examples 

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html

c:\Users\lucas\go\bin\build-lambda-zip.exe -o detectFaceHandler.zip detectFaceHandler

package main

import (
    "context"
    "fmt"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func lambda_handler(ctx context.Context, s3Event events.S3Event) {
    for _, record := range s3Event.Records {
        s3 := record.S3
        fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
    }
}

func main() {
    lambda.Start(lambda_handler)
}
GOOS=linux go build detectFaceHandler.go
go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-lambda-go/events
zip detectFaceHandler.zip detectFaceHandler
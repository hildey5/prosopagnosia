package prosopagnosia

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// uploadPhoto - This function is uploading a photo to a s3 bucket.
func uploadPhoto(bucket string, filename string) error {
	// If there is a error with opening the file this will output the error.
	file, err := os.Open(filename)
	if err != nil {
		printErrorf("Unable to open file %q, %v", err)
		return err
	}

	defer file.Close()

	// Here it creats a new session and also says the region where the bucket was created for security purposes.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	uploader := s3manager.NewUploader(sess)

	// This is where is uploads the file to the s3 bucket.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		Key: aws.String(filename),

		Body: file,

		//If there is an error when uplading the file to the s3 bucket, this will output a error.
	})
	if err != nil {
		// Print the error and exit.
		printErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		return err
	}

	// If the file is successfully uploaded to the s3 bucket this will say that it was successfull.
	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
	return nil
}

func printErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

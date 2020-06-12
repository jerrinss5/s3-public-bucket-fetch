package main

import (
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
)

func main() {
	regionPtr := flag.String("r", "", "Region where the bucket is located")
	bucketPtr := flag.String("b", "", "Bucket name")
	keyPtr := flag.String("k", "", "key name (name of the file)")
	fileNamePtr := flag.String("f", "", "filename for the downloaded file")

	flag.Parse()

	if *regionPtr == "" || *bucketPtr == "" || *keyPtr == "" || *fileNamePtr == "" {
		fmt.Println("You must supply a region name, bucket name, key name and the file name.")
		os.Exit(0)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(*regionPtr),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(*fileNamePtr)
	if err != nil {
		fmt.Printf("failed to create file %q, %v", "downloaded", err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(*bucketPtr),
		Key:    aws.String(*keyPtr),
	})
	if err != nil {
		fmt.Printf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
}

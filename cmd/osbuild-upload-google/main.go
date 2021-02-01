package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/osbuild/osbuild-composer/internal/upload/google"
)

func checkStringNotEmpty(variable string, errorMessage string) {
	if variable == "" {
		fmt.Fprintln(os.Stderr, errorMessage)
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	var credentialsPath string
	var sourcePath string
	var bucket string
	var destPath string

	flag.StringVar(&credentialsPath, "creds-path", "", "Path to service account credentials (mandatory)")
	flag.StringVar(&sourcePath, "source", "", "File to upload (mandatory)")
	flag.StringVar(&bucket, "bucket", "", "Bucket name (mandatory)")
	flag.StringVar(&destPath, "dest", "", "Target filename in the bucket (mandatory)")

	flag.Parse()

	checkStringNotEmpty(credentialsPath, "You need to specify the path to the service account credentials json file")
	checkStringNotEmpty(sourcePath, "You need to specify the file to copy")
	checkStringNotEmpty(bucket, "You need to specify the bucket name")
	checkStringNotEmpty(destPath, "You need to specify the filename in the bucket")

	creds, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		fmt.Println("ERROR: err", err)
		return
	}


	g, err := google.New(creds)
	if err != nil {
		fmt.Println("ERROR: err", err)
		return
	}

	err = g.Upload(sourcePath, destPath, bucket)
	if err != nil {
		fmt.Println("ERROR: err", err)
		return
	}
}

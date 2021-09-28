package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if len(os.Getenv("AWS_EXECUTION_ENV")) == 0 {
		commandLineHandler()
	} else {
		lambda.Start(cloudHandler)
	}
}

func cloudHandler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3

		log.Println("*** STARTING LAMBDA HANDLER ***")
		log.Println(s3.Bucket)

		log.Printf("*** ENDING LAMBDA HANDLER ***")
	}
}

func commandLineHandler() {
	var err error
	var fileName string

	flag.StringVar(&fileName, "i", "", "input file")
	flag.Parse()

	if len(fileName) == 0 {
		log.Println("RETINA")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal(err.Error())
	}
}

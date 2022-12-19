package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) error {
	s := os.Getenv("PERMIT")
	log.Printf("PERMIT: %s", s)
	return nil
}

func main() {
	lambda.Start(handler)
}

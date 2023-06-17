package main

import (
	"context"
	"log"

	"example.com/go_functions"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	ctx := context.Background()
	tweet := go_functions.Tweet

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", tweet); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

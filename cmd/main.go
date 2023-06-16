package main

import (
	"context"
	"log"

	"example.com/go_functions"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	ctx := context.Background()

	hello := go_functions.HelloWorld

	// 実行する関数の登録
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", hello); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

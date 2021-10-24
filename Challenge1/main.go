package main

import (
	"context"
	"fmt"
	"log"
	"cloud.google.com/go/translate"
)

func main(){
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
		// TODO: handle error.
	}
	ds, err := client.DetectLanguage(ctx, []string{"Today is Monday"})
	if err != nil {
		log.Fatal(err)
		// TODO: handle error.
	}
	fmt.Println(ds)
	fmt.Println()


	
}
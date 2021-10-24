package main


import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func translateTextWithModel(targetLanguage, text, model string) (string, error) {
	// targetLanguage = "ja"
	// text = "The Go Gopher is cute"
	// model = "nmt"

	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
			return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
			return "", fmt.Errorf("translate.NewClient: %v", err)
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
			Model: model, // Either "nmt" or "base".
	})
	if err != nil {
			return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
			return "", nil
	}
	return resp[0].Text, nil
}

func main(){
	fmt.Println(translateTextWithModel("ja","tai","nmt"))
}
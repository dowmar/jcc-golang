package transla

import (
	"context"
	"log"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func main() {
	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx)
	if err != nil {
		log.Fatal(err)
		// TODO: Handle error.
	}
	defer c.Close()

	req := &translatepb.TranslateTextRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/translate/v3#TranslateTextRequest.
	}
	resp, err := c.TranslateText(ctx, req)
	if err != nil {
		log.Fatal(err)
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

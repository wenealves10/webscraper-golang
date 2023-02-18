package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://codedev.digital"),
		chromedp.Title(&title),
	)
	if err != nil {
		fmt.Println("Error executing Chromedp:", err)
		return
	}
	fmt.Println("Page title:", title)

}

package domain

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func Navigate(domain string) {

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// fmt.Println(domain)
	if ctx == nil {
		fmt.Println("Context error")
		return
	}

	// navigate to a page
	if err := chromedp.Run(ctx,
		chromedp.Navigate(domain),
		chromedp.ActionFunc(cookies),
	); err != nil {
		log.Fatal(err)
	}

}

func cookies(ctx context.Context) error {
	cookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		return err
	}

	for i, cookie := range cookies {
		fmt.Printf("Cookie  %d: %v\n", i, cookie.Name)
		// TODO write in file

		if verifyJWT(cookie) {
			return nil
		}
	}

	return nil
}

func verifyJWT(cookie *network.Cookie) bool {
	if strings.Contains(cookie.Name, "JWT") || strings.Contains(cookie.Name, "jwt") {
		fmt.Println(cookie.Domain, "JWT")
		// TODO write in file
		return true
	}
	return false
}

func writeFile() {
	// TODO write results in file
}

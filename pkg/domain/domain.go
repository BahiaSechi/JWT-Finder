package domain

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Open a browser
func Navigate(domain string) {

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	if ctx == nil {
		fmt.Println("Context error")
		return
	}

	// navigate to a page
	if err := chromedp.Run(ctx,
		chromedp.Navigate(domain),
		chromedp.ActionFunc(cookies),
		// Check Authrozation header and print it
		chromedp.ActionFunc(func(ctx context.Context) error {
			resp, err := http.Get(domain)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Authorization : " + resp.Header.Get("Authorization"))
			//fmt.Println("Content-Security-Policy : " + resp.Header.Get("Content-Security-Policy"))
			//fmt.Println("Last-Modified : " + resp.Header.Get("Last-Modified"))
			//fmt.Println("Date : " + resp.Header.Get("Date"))

			return nil
		}),
	); err != nil {
		log.Fatal(err)
	}

}

// Retrieve cookies
func cookies(ctx context.Context) error {

	cookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		return err
	}

	// Display every cookie
	for i, cookie := range cookies {
		fmt.Printf("Cookie  %d: %v\n", i, cookie.Name)
		// TODO write in file

		// Check if one of the cookies is named JWT
		if verifyJWT(cookie) {
			return nil
		}
	}

	return nil
}

// Check if the cookie is named JWT or jwt
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

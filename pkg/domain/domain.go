package domain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Open a browser
func Navigate(domain string) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	if ctx == nil {
		fmt.Println("Context error")
		return
	}
	chromedp.ListenTarget(ctx, func(event interface{}) {
		switch responseReceivedEvent := event.(type) {
		case *network.EventResponseReceived:
			response := responseReceivedEvent.Response

			if auth, ok := response.Headers["Authorization"]; ok {
				log.Printf("%s", auth)
			}
			if auth, ok := response.Headers["authorization"]; ok {
				log.Printf("%s", auth)
			}
		}
	})

	// navigate to a page
	if err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(domain),
		chromedp.ActionFunc(cookies),
		/* Crawler
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			response, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			fmt.Println(response)
			//log.Printf("%s", response)

			return err
		}),*/
	); err != nil {
		fmt.Println("\033[31m", err, "\033[0m")
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

		// Check if one of the cookies is named JWT
		if verifyJWT(cookie) {
			return nil
		}
	}

	fmt.Println("\033[36m", "Number of cookies :", len(cookies), "\033[0m")

	return nil
}

// Check if the cookie is named JWT or jwt
func verifyJWT(cookie *network.Cookie) bool {
	if strings.Contains(strings.ToLower(cookie.Name), "jwt") {
		fmt.Println("\033[32m", cookie.Domain, "JWT", "\033[0m")
		// TODO write in file
		return true
	}
	return false
}

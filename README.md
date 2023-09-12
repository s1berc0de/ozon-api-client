# Ozon Seller API Client
A Ozon Seller API client written in Golang

[![Coverage Status](https://coveralls.io/repos/github/diPhantxm/ozon-api-client/badge.svg)](https://coveralls.io/github/diPhantxm/ozon-api-client)
![example workflow](https://github.com/diPhantxm/ozon-api-client/actions/workflows/tests.yml/badge.svg)

[Ozon](https://ozon.ru) is a marketplace for small and medium enterprises to launch and grow their businesses in Russia.

Read full [documentation](https://docs.ozon.ru/api/seller/en/#tag/Introduction)

You can check [list of supported endpoints](ENDPOINTS.md)

## How to start
### API
Get Client-Id and Api-Key in your seller profile [here](https://seller.ozon.ru/app/settings/api-keys?locale=en)

Just add dependency to your project and you're ready to go.
```bash
go get github.com/diphantxm/ozon-api-client
```
A simple example on how to use this library:
```Golang
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func main() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	client := ozon.NewClient("my-client-id", "my-api-key")

	// Send request with parameters
	resp, err := client.Products().GetProductDetails(&ozon.GetProductDetailsParams{
		ProductId: 123456789,
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("error when getting product details: %s", err)
	}

	// Do some stuff
	for _, d := range resp.Result.Barcodes {
		fmt.Printf("Barcode %s\n", d)
	}
}
```

### Notifications
Ozon can send push-notifications to your REST server. There is an implementation of REST server that handles notifications in this library.

[Official documentation](https://docs.ozon.ru/api/seller/en/#tag/push_intro)

How to use:
```Golang
package main

import (
	"log"

	"github.com/diphantxm/ozon-api-client/ozon/notifications"
)

func main() {
	// Create server
	port := 5000
	server := notifications.NewNotificationServer(port)

	// Register handlers passing message type and handler itself
	server.Register(notifications.ChatClosedType, func(req interface{}) error {
		notification := req.(*notifications.ChatClosed)

		// Do something with the notification here...
		log.Printf("chat %s has been closed\n", notification.ChatId)

		return nil
	})

	// Run server
	if err := server.Run(); err != nil {
		log.Printf("error while running notification server: %s", err)
	}
}
```

## Contribution
If you need some endpoints ASAP, create an issue and list all the endpoints. I will add them to library soon.

Or you can implement them and contribute to the project. Contribution to the project is welcome. 

## Endpoints
1. /v2/category/tree ✅
2. /v3/category/attribute ✅
3. /v2/category/attribute/values ✅
4. /v2/product/import ✅
5. /v1/product/import/info ✅
6. /v1/product/import-by-sku ✅
7. /v1/product/attributes/update ✅
8. /v1/product/pictures/import ✅
9. /v1/product/pictures/info ✅
10. /v2/product/list ✅
11. /v2/product/info ✅
12. /v1/product/rating-by-sku ✅
13. /v2/product/info/list ✅
14. /v3/products/info/attributes
15. /v1/product/info/description
16. /v4/product/info/limit
17. /v1/product/update/offer-id
18. /v1/product/archive
19. /v1/product/unarchive
20. /v2/products/delete
21. /v1/products/geo-restrictions-catalog-by-filter
22. /v1/product/upload_digital_codes
23. /v1/product/upload_digital_codes/info
24. /v1/product/info/subscription

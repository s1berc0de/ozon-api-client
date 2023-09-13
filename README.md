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
### Атрибуты и характеристики Ozon
1. /v2/category/tree ✅
2. /v3/category/attribute ✅
3. /v2/category/attribute/values ✅
### Загрузка и обновление товаров
1. /v2/product/import ✅
2. /v1/product/import/info ✅
3. /v1/product/import-by-sku ✅
4. /v1/product/attributes/update ✅
5. /v1/product/pictures/import ✅
6. /v1/product/pictures/info ✅
7. /v2/product/list ✅
8. /v2/product/info ✅
9. /v1/product/rating-by-sku ✅
10. /v2/product/info/list ✅
11. /v3/products/info/attributes ✅
12. /v1/product/info/description ✅
13. /v4/product/info/limit ✅
14. /v1/product/update/offer-id ✅
15. /v1/product/archive ✅
16. /v1/product/unarchive ✅
17. /v2/products/delete ✅
18. /v1/products/geo-restrictions-catalog-by-filter ✅
19. /v1/product/upload_digital_codes ✅
20. /v1/product/upload_digital_codes/info ✅
21. /v1/product/info/subscription ✅
### Цены и остатки товаров
1. /v1/product/import/stocks ✅
2. /v2/products/stocks ✅
3. /v3/product/info/stocks ✅
4. /v1/product/info/stocks-by-warehouse/fbs ✅
5. /v1/product/import/prices ✅
6. /v4/product/info/prices ✅
7. /v1/product/info/discounted ✅
8. /v1/product/update/discount ✅
### Акции
1. /v1/actions [GET]
2. /v1/actions/candidates
3. /v1/actions/products
4. /v1/actions/products/activate
5. /v1/actions/products/deactivate
6. /v1/actions/hotsales/list
7. /v1/actions/hotsales/products
8. /v1/actions/hotsales/activate
9. /v1/actions/hotsales/deactivate
10. /v1/actions/discounts-task/list
11. /v1/actions/discounts-task/approve
12. /v1/actions/discounts-task/decline
### Стратегии ценообразования
1. /v1/pricing-strategy/competitors/list
2. /v1/pricing-strategy/list
3. /v1/pricing-strategy/create
4. /v1/pricing-strategy/info
5. /v1/pricing-strategy/update
6. /v1/pricing-strategy/products/add
7. /v1/pricing-strategy/strategy-ids-by-product-ids
8. /v1/pricing-strategy/products/list
9. /v1/pricing-strategy/product/info
10. /v1/pricing-strategy/products/delete
11. /v1/pricing-strategy/status
12. /v1/pricing-strategy/delete
### Сертификаты брендов
1. /v1/brand/company-certification/list
### Сертификаты качества
1. /v1/product/certificate/accordance-types [GET]
2. /v2/product/certificate/accordance-types/list [GET]
3. /v1/product/certificate/types
4. /v1/product/certification/list
5. /v1/product/certificate/create
6. /v1/product/certificate/bind
7. /v1/product/certificate/delete
8. /v1/product/certificate/info
9. /v1/product/certificate/list
10. /v1/product/certificate/product_status/list
11. /v1/product/certificate/products/list
12. /v1/product/certificate/unbind
13. /v1/product/certificate/rejection_reasons/list
14. /v1/product/certificate/status/list
### Склады
1. /v1/warehouse/list
2. /v1/delivery-method/list
### Полигоны
1. /v1/polygon/create
2. /v1/polygon/bind
### Схема FBO
1. /v2/posting/fbo/list
2. /v2/posting/fbo/get
3. /v1/supply-order/list
4. /v1/supply-order/get
5. /v1/supply-order/items
6. /v1/supplier/available_warehouses [GET]
### Управление кодами маркировки для FBS/rFBS
1. /v4/fbs/posting/product/exemplar/validate
2. /v4/fbs/posting/product/exemplar/set
3. /v4/fbs/posting/product/exemplar/status
4. /v4/posting/fbs/ship
### Схемы FBS и rFBS
1. /v3/posting/fbs/unfulfilled/list
2. /v3/posting/fbs/list
3. /v3/posting/fbs/get
4. /v2/posting/fbs/get-by-barcode
5. /v2/posting/fbs/act/get-barcode
6. /v2/posting/fbs/act/get-barcode/text
7. /v2/posting/fbs/product/country/list
8. /v2/posting/fbs/product/country/set
9. /v3/posting/multiboxqty/set
10. /v1/posting/fbs/restrictions
11. /v3/posting/fbs/ship
12. /v3/posting/fbs/ship/package
13. /v2/posting/fbs/act/create
14. /v2/posting/fbs/act/check-status
15. /v1/posting/carriage-available/list
16. /v2/posting/fbs/act/get-pdf
17. /v2/posting/fbs/digital/act/check-status
18. /v2/posting/fbs/digital/act/get-pdf
19. /v2/posting/fbs/package-label
20. /v1/posting/fbs/package-label/create
21. /v1/posting/fbs/package-label/get
22. /v2/posting/fbs/act/get-container-labels
23. /v2/posting/fbs/arbitration
24. /v2/posting/fbs/awaiting-delivery
25. /v1/posting/fbs/cancel-reason
26. /v2/posting/fbs/cancel-reason/list
27. /v2/posting/fbs/cancel
28. /v2/posting/fbs/product/change
29. /v2/posting/fbs/product/cancel
30. /v2/posting/fbs/act/list
31. /v2/posting/fbs/digital/act/document-sign
32. /v2/posting/fbs/act/get-postings
33. /v2/fbs/posting/delivering
34. /v2/fbs/posting/tracking-number/set
35. /v2/fbs/posting/last-mile
36. /v2/fbs/posting/delivered
37. /v2/fbs/posting/sent-by-seller
38. /v1/posting/fbs/timeslot/change-restrictions
39. /v1/posting/fbs/timeslot/set
40. /v1/posting/global/etgb
### Возвраты товаров
1. /v3/returns/company/fbo
2. /v3/returns/company/fbs
### Отмены заказов
1. /v1/conditional-cancellation/get
2. /v1/conditional-cancellation/list
3. /v1/conditional-cancellation/approve
4. /v1/conditional-cancellation/reject
### Чаты с покупателями
1. /v2/chat/list
2. /v1/chat/send/message
3. /v1/chat/send/file
4. /v2/chat/history
5. /v1/chat/updates
6. /v1/chat/start
7. /v2/chat/read
### Накладные
1. /v1/invoice/create-or-update
2. /v1/invoice/get
3. /v1/invoice/delete
### Отчёты
1. /v1/report/info
2. /v1/report/list
3. /v1/report/products/create
4. /v1/report/stock/create
5. /v1/report/products/movement/create
6. /v1/report/returns/create
7. /v1/report/postings/create
8. /v1/finance/cash-flow-statement/list
9. /v1/report/discounted/create
10. /v1/report/discounted/info
11. /v1/report/discounted/list
### Аналитические отчёты
1. /v1/analytics/data
2. /v2/analytics/stock_on_warehouses
### Финансовые отчёты
1. /v1/finance/realization
2. /v3/finance/transaction/list
3. /v3/finance/transaction/totals
### Рейтинг продавца
1. /v1/rating/summary
2. /v1/rating/history
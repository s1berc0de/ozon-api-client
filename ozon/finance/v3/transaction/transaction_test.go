package transaction_test

import (
	"bytes"
	"context"
	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/finance/v3/transaction"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestList_Success(t *testing.T) {
	c := transaction.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/finance/transaction/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"date":{"from":"2021-11-01T00:00:00Z","to":"2021-11-02T00:00:00Z"},"operation_type":[],"posting_number":"","transaction_type":"all"},"page":1,"page_size":1000}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							  "result": {
								"operations": [
								  {
									"operation_id": 11401182187840,
									"operation_type": "MarketplaceMarketingActionCostOperation",
									"operation_date": "2021-11-01 00:00:00",
									"operation_type_name": "Услуги продвижения товаров",
									"delivery_charge": 0,
									"return_delivery_charge": 0,
									"accruals_for_sale": 0,
									"sale_commission": 0,
									"amount": -6.46,
									"type": "services",
									"posting": {
									  "delivery_schema": "",
									  "order_date": "",
									  "posting_number": "",
									  "warehouse_id": 0
									},
									"items": [],
									"services": []
								  }
								],
								"page_count": 1,
								"row_count": 355
							}}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/finance/transaction",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &transaction.ListRequest{
		Filter: transaction.ListRequestFilter{
			Date: transaction.ListRequestFilterDate{
				From: time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2021, 11, 2, 0, 0, 0, 0, time.UTC),
			},
			OperationType:   []string{},
			PostingNumber:   "",
			TransactionType: "all",
		},
		Page:     1,
		PageSize: 1000,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &transaction.ListResponse{
		Result: transaction.ListResponseResult{
			Operations: []transaction.ListResponseResultOperation{
				{
					OperationID:          11401182187840,
					OperationType:        "MarketplaceMarketingActionCostOperation",
					OperationDate:        "2021-11-01 00:00:00",
					OperationTypeName:    "Услуги продвижения товаров",
					DeliveryCharge:       0,
					ReturnDeliveryCharge: 0,
					AccrualsForSale:      0,
					SaleCommission:       0,
					Amount:               -6.46,
					Type:                 "services",
					Posting: transaction.ListResponseResultOperationPosting{
						DeliverySchema: "",
						OrderDate:      "",
						PostingNumber:  "",
						WarehouseID:    0,
					},
					Items: []transaction.ListResponseResultOperationItem{},
				},
			},
			PageCount: 1,
			RowCount:  355,
		},
	}, resp)
}

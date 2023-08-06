package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	v1 "github.com/diphantxm/ozon-api-client/ozon/product/v1"
	"github.com/stretchr/testify/require"
)

func TestImportBySKU_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import-by-sku", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"items":[{"sku":298789742,"name":"string","offer_id":"91132","currency_code":"RUB","old_price":"2590","price":"2300","premium_price":"2200","vat":"0.1"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  	"result": {
								"task_id": 176594213,
								"unmatched_sku_list": [ ]
						  	}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.ImportBySKU(context.Background(), &v1.ImportBySKURequest{
		Items: []v1.ImportBySKURequestItem{
			{
				SKU:          298789742,
				Name:         "string",
				OfferID:      "91132",
				OldPrice:     "2590",
				Price:        "2300",
				PremiumPrice: "2200",
				Vat:          "0.1",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.ImportBySKUResponse{
		Result: v1.ImportBySKUResponseResult{
			TaskID:           176594213,
			UnmatchedSKUList: make([]int64, 0),
		},
	}, resp)
}

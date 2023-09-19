package update_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/ozon/product/v1/update"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestOfferID_Success(t *testing.T) {
	c := update.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/update/offer-id", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"update_offer_id":[{"new_offer_id":"string","offer_id":"string"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"errors": [
								{
									"message": "string",
									"offer_id": "string"
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/update",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.OfferID(context.Background(), &update.OfferIDRequest{
		UpdateOfferID: []update.OfferIDRequestUpdateOfferIDItem{
			{
				NewOfferID: "string",
				OfferID:    "string",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &update.OfferIDResponse{
		Errors: []update.OfferIDResponseError{
			{
				Message: "string",
				OfferID: "string",
			},
		},
	}, resp)
}

func TestDiscount_Success(t *testing.T) {
	c := update.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/update/discount", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"discount":0,"product_id":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/update",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Discount(context.Background(), &update.DiscountRequest{
		Discount:  0,
		ProductID: 0,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &update.DiscountResponse{
		Result: true,
	}, resp)
}

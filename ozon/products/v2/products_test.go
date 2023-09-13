package v2_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	v2 "github.com/diphantxm/ozon-api-client/ozon/products/v2"
	"github.com/stretchr/testify/require"
)

func TestDelete_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/products/delete", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"products":[{"offer_id":"033"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"status": [
								{
									"offer_id": "033",
									"is_deleted": true,
									"error": ""
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/products",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Delete(context.Background(), &v2.DeleteRequest{
		Products: []v2.DeleteRequestProduct{
			{
				OfferID: "033",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.DeleteResponse{
		Status: []v2.DeleteResponseStatus{
			{
				OfferID:   "033",
				IsDeleted: true,
				Error:     "",
			},
		},
	}, resp)
}

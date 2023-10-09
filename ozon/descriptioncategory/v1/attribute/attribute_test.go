package attribute_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1/attribute"
	"github.com/stretchr/testify/require"
)

func TestValues_Success(t *testing.T) {
	c := attribute.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/description-category/attribute/values", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"attribute_id":0,"category_id":0,"language":"DEFAULT","last_value_id":0,"limit":0,"type_id":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"has_next": true,
							"result": [
								{
									"id": 0,
									"info": "string",
									"picture": "string",
									"value": "string"
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/description-category/attribute",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Values(context.Background(), &attribute.ValuesRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &attribute.ValuesResponse{
		Result: []attribute.ValuesResponseResult{
			{
				ID:      0,
				Info:    "string",
				Picture: "string",
				Value:   "string",
			},
		},
		HasNext: true,
	}, resp)
}

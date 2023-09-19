package attribute_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/category/v2/attribute"
	"github.com/stretchr/testify/require"
)

func TestValues_Success(t *testing.T) {
	c := attribute.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/category/attribute/values", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"attribute_id":10096,"category_id":17028968,"language":"DEFAULT","last_value_id":0,"limit":3}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"id": 61571,
									"value": "белый",
									"info": "",
									"picture": ""
								},
								{
									"id": 61572,
									"value": "прозрачный",
									"info": "",
									"picture": ""
								},
								{
									"id": 61573,
									"value": "бежевый",
									"info": "",
									"picture": ""
								}
							],
							"has_next": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/category/attribute",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Values(context.Background(), &attribute.ValuesRequest{
		AttributeId: 10096,
		CategoryId:  17028968,
		LastValueId: 0,
		Limit:       3,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &attribute.ValuesResponse{
		Result: []attribute.ValuesResponseResult{
			{
				ID:      61571,
				Value:   "белый",
				Info:    "",
				Picture: "",
			},
			{
				ID:      61572,
				Value:   "прозрачный",
				Info:    "",
				Picture: "",
			},
			{
				ID:      61573,
				Value:   "бежевый",
				Info:    "",
				Picture: "",
			},
		},
		HasNext: true,
	}, resp)
}

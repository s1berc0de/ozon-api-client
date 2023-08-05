package attribute_test

import (
	"bytes"
	"context"
	"github.com/diphantxm/ozon-api-client/internal/auth"
	attribute2 "github.com/diphantxm/ozon-api-client/internal/ozon/category/v2/attribute"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestValues_Success(t *testing.T) {
	c := attribute2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(req *http.Request) *http.Response {
					require.Equal(t, req.Header.Get(auth.APIKeyHeader), test.ApiKey)
					require.Equal(t, req.Header.Get(auth.ClientIDHeader), test.ClientID)

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
		"any_uri",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Values(context.Background(), &attribute2.ValuesRequest{
		AttributeId: 10096,
		CategoryId:  17028968,
		LastValueId: 0,
		Limit:       3,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &attribute2.ValuesResponse{
		Result: []attribute2.ValuesResponseResult{
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

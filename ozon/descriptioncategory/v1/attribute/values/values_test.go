package values_test

import (
	"bytes"
	"context"
	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/descriptioncategory/v1/attribute/values"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestSearch_Success(t *testing.T) {
	c := values.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/description-category/attribute/values/search", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"attribute_id":0,"description_category_id":0,"limit":0,"type_id":0,"value":""}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
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
		"https://api-seller.ozon.ru/v1/description-category/attribute/values",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Search(context.Background(), &values.SearchRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &values.SearchResponse{
		Result: []values.SearchResponseResult{
			{
				ID:      0,
				Info:    "string",
				Picture: "string",
				Value:   "string",
			},
		},
	}, resp)
}

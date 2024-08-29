package v2_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	v2 "github.com/andmetoo/ozon-api-client/ozon/category/v2"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestTree_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/category/tree", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"category_id":17034410,"language":"DEFAULT"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
				  				{
									"category_id": 17034410,
									"title": "Развивающие игрушки",
									"children": [ ]
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/category",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Tree(context.Background(), &v2.TreeRequest{
		CategoryId: 17034410,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.TreeResponse{
		Result: []v2.TreeResponseResult{
			{
				CategoryId: 17034410,
				Children:   make([]v2.TreeResponseResult, 0),
				Title:      "Развивающие игрушки",
			},
		},
	}, resp)
}

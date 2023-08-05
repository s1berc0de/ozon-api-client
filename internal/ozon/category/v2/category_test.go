package v2_test

import (
	"bytes"
	"context"
	"github.com/diphantxm/ozon-api-client/internal/auth"
	v22 "github.com/diphantxm/ozon-api-client/internal/ozon/category/v2"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestTree_Success(t *testing.T) {
	c := v22.New(
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
									"category_id": 17034410,
									"title": "Развивающие игрушки",
									"children": []
								}
							]
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

	resp, httpResp, err := c.Tree(context.Background(), &v22.TreeRequest{
		CategoryId: 17034410,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v22.TreeResponse{
		Result: []v22.TreeResponseResult{
			{
				CategoryId: 17034410,
				Children:   make([]v22.TreeResponseResult, 0),
				Title:      "Развивающие игрушки",
			},
		},
	}, resp)
}

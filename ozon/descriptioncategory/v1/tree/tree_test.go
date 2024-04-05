package tree_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1/tree"
	"github.com/stretchr/testify/require"
)

func TestTree_Success(t *testing.T) {
	c := tree.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/description-category/tree", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"language":"DEFAULT"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [{
							  "description_category_id": 0,
							  "category_name": "string",
							  "disabled": false,
							  "children": [
								{
								  "description_category_id": 0,
								  "category_name": "string",
								  "disabled": false,
								  "children": [
									{
									  "type_name": "sting",
									  "type_id": 0,
									  "disabled": false,
									  "children": []
									}
								  ]
								}
							  ]
							}
						  ]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/description-category/tree",
	)

	resp, httpResp, err := c.Tree(context.Background(), &tree.TreeRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &tree.TreeResponse{
		Result: []tree.TreeResult{
			{
				DescriptionCategoryID: ptr(int64(0)),
				CategoryName:          ptr("string"),
				Disabled:              false,
				Children: []*tree.TreeResult{
					{
						DescriptionCategoryID: ptr(int64(0)),
						CategoryName:          ptr("string"),
						Disabled:              false,
						Children: []*tree.TreeResult{
							{
								TypeName: ptr("sting"),
								TypeID:   ptr(int64(0)),
								Disabled: false,
								Children: make([]*tree.TreeResult, 0),
							},
						},
					},
				},
			},
		},
	}, resp)

	require.NotNil(t, c)
}

func ptr[T any](v T) *T {
	return &v
}

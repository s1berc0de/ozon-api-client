package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1"
	"github.com/stretchr/testify/require"
)

func TestTree_Success(t *testing.T) {
	c := v1.New(
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
		"https://api-seller.ozon.ru/v1/description-category",
	)

	resp, httpResp, err := c.Tree(context.Background(), &v1.TreeRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.TreeResponse{
		Result: []*v1.TreeResult{
			{
				DescriptionCategoryID: ptr(int64(0)),
				CategoryName:          ptr("string"),
				Disabled:              false,
				Children: []*v1.TreeResult{
					{
						DescriptionCategoryID: ptr(int64(0)),
						CategoryName:          ptr("string"),
						Disabled:              false,
						Children: []*v1.TreeResult{
							{
								TypeName: ptr("sting"),
								TypeID:   ptr(int64(0)),
								Disabled: false,
								Children: make([]*v1.TreeResult, 0),
							},
						},
					},
				},
			},
		},
	}, resp)

	require.NotNil(t, c)
}

func TestAttribute_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/description-category/attribute", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"description_category_id":0,"language":"DEFAULT","type_id":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
								"result": [
									{
									  "category_dependent": true,
									  "description": "string",
									  "dictionary_id": 0,
									  "group_id": 0,
									  "group_name": "string",
									  "id": 0,
									  "is_aspect": true,
									  "is_collection": true,
									  "is_required": true,
									  "name": "string",
									  "type": "string",
									  "attribute_complex_id": 0,
									  "max_value_count": 0
									}
								]
							  }`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/description-category",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Attribute(context.Background(), &v1.AttributeRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.AttributeResponse{
		Result: []v1.AttributeResponseResult{
			{
				CategoryDependent:  true,
				Description:        "string",
				DictionaryID:       0,
				GroupID:            0,
				GroupName:          "string",
				ID:                 0,
				IsAspect:           true,
				IsCollection:       true,
				IsRequired:         true,
				Name:               "string",
				Type:               "string",
				AttributeComplexID: 0,
				MaxValueCount:      0,
			},
		},
	}, resp)
}

func ptr[T any](v T) *T {
	return &v
}

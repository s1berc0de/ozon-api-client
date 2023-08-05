package v3_test

import (
	"bytes"
	"context"
	"github.com/diphantxm/ozon-api-client/auth"
	v3 "github.com/diphantxm/ozon-api-client/ozon/category/v3"
	"github.com/diphantxm/ozon-api-client/test"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestAttribute_Success(t *testing.T) {
	c := v3.New(
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
									"attributes": [
								  		{
											"id": 85,
											"name": "Бренд",
											"description": "Укажите наименование бренда, под которым произведен товар. Если товар не имеет бренда, используйте значение \"Нет бренда\"",
											"type": "String",
											"is_collection": false,
											"is_required": true,
											"group_id": 0,
											"group_name": "",
											"dictionary_id": 28732849
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
		"any_uri",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Attribute(context.Background(), &v3.AttributeRequest{
		CategoryId: []int64{17034410},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v3.AttributeResponse{
		Result: []v3.AttributeResponseResult{
			{
				CategoryId: 17034410,
				Attributes: []v3.AttributeResponseResultAttribute{
					{
						ID:           85,
						Name:         "Бренд",
						Description:  "Укажите наименование бренда, под которым произведен товар. Если товар не имеет бренда, используйте значение \"Нет бренда\"",
						Type:         "String",
						IsCollection: false,
						IsRequired:   true,
						GroupId:      0,
						GroupName:    "",
						DictionaryId: 28732849,
					},
				},
			},
		},
	}, resp)
}

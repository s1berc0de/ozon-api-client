package attributes_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/product/v1/attributes"
	"github.com/stretchr/testify/require"
)

func TestUpdate_Success(t *testing.T) {
	c := attributes.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/attributes/update", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"items":[{"attributes":[{"complex_id":0,"id":0,"values":[{"dictionary_value_id":0,"value":"string"}]}],"offer_id":"string"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"task_id": 0
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/attributes",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Update(context.Background(), &attributes.UpdateRequest{
		Items: []attributes.UpdateRequestItem{
			{
				Attributes: []attributes.UpdateRequestItemAttribute{
					{
						Values: []attributes.UpdateRequestItemAttributeValue{
							{
								Value: "string",
							},
						},
					},
				},
				OfferID: "string",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &attributes.UpdateResponse{
		TaskID: 0,
	}, resp)
}

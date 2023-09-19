package uploaddigitalcodes_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/product/v1/uploaddigitalcodes"
	"github.com/stretchr/testify/require"
)

func TestInfo_Success(t *testing.T) {
	c := uploaddigitalcodes.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/upload_digital_codes/info", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"task_id":178574231}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"status": "imported"
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/upload_digital_codes",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Info(context.Background(), &uploaddigitalcodes.InfoRequest{
		TaskID: 178574231,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &uploaddigitalcodes.InfoResponse{
		Result: uploaddigitalcodes.InfoResponseResult{
			Status: uploaddigitalcodes.InfoResponseResultStatusImported,
		},
	}, resp)
}

package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/diphantxm/ozon-api-client/ozon/product/v4/info"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestLimit_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v4/product/info/limit", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"daily_create": {
								"limit": 0,
								"reset_at": "2019-08-24T14:15:22Z",
								"usage": 0
							},
							"daily_update": {
								"limit": 0,
								"reset_at": "2019-08-24T14:15:22Z",
								"usage": 0
							},
							"total": {
								"limit": 0,
								"usage": 0
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v4/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Limit(context.Background(), &info.LimitRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.LimitResponse{
		DailyCreate: info.LimitResponseDailyCreate{
			Limit:   0,
			ResetAt: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
			Usage:   0,
		},
		DailyUpdate: info.LimitResponseDailyUpdate{
			Limit:   0,
			ResetAt: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
			Usage:   0,
		},
		Total: info.LimitResponseTotal{
			Limit: 0,
			Usage: 0,
		},
	}, resp)
}

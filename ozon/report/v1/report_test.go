package v1_test

import (
	"bytes"
	"context"
	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	v1 "github.com/s1berc0de/ozon-api-client/ozon/report/v1"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestInfo_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/report/info", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"code":"REPORT_seller_products_924336_1720170405_a9ea2f27-a473-4b13-99f9-d0cfcb5b1a69"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(bytes.NewBufferString(`{"result":{"code": "REPORT_seller_products_924336_1720170405_a9ea2f27","status": "success","error": "","file": "https://cdn1.ozone.ru/s3/item-picture-6/f3/ce/f4ceae54b323213d3e61e59c323bd8e5.csv","report_type":"seller_products","params":{},"created_at":"2021-11-25T14:54:55.688260Z"}}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/report",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Info(context.Background(), &v1.InfoRequest{
		Code: "REPORT_seller_products_924336_1720170405_a9ea2f27-a473-4b13-99f9-d0cfcb5b1a69",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.InfoResponse{
		Result: v1.InfoResponseResult{
			Code:       "REPORT_seller_products_924336_1720170405_a9ea2f27",
			Status:     "success",
			Error:      "",
			File:       "https://cdn1.ozone.ru/s3/item-picture-6/f3/ce/f4ceae54b323213d3e61e59c323bd8e5.csv",
			ReportType: "seller_products",
			CreatedAt:  time.Date(2021, 11, 25, 14, 54, 55, 688260000, time.UTC),
		},
	}, resp)
}

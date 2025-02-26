package postings_test

import (
	"bytes"
	"context"
	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	"github.com/s1berc0de/ozon-api-client/ozon/report/v1/postings"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestCreate_Success(t *testing.T) {
	c := postings.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/report/postings/create", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"processed_at_from":"2021-09-02T17:10:54.861Z","processed_at_to":"2021-11-02T17:10:54.861Z","delivery_schema":["fbo"],"sku":[],"cancel_reason_id":[],"offer_id":"offer1","status_alias":[],"statuses":[],"title":""},"language":"DEFAULT"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(bytes.NewBufferString(`{"result":{"code":"REPORT_seller_postings_514893_1722847571_32a3508c-6b53-408c-a212-6c97138d23ed"}}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/report/postings",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Create(context.Background(), &postings.CreateRequest{
		Filter: postings.CreateRequestFilter{
			ProcessedAtFrom: time.Date(2021, 9, 2, 17, 10, 54, 861000000, time.UTC),
			ProcessedAtTo:   time.Date(2021, 11, 2, 17, 10, 54, 861000000, time.UTC),
			DeliverySchema:  []string{"fbo"},
			Sku:             []int64{},
			CancelReasonId:  []int64{},
			OfferId:         "offer1",
			StatusAlias:     []string{},
			Statuses:        []int64{},
			Title:           "",
		},
		Language: "DEFAULT",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &postings.CreateResponse{Result: postings.CreateResponseResult{Code: "REPORT_seller_postings_514893_1722847571_32a3508c-6b53-408c-a212-6c97138d23ed"}}, resp)
}

package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	v1 "github.com/diphantxm/ozon-api-client/ozon/rating/v1"
	"github.com/stretchr/testify/require"
)

func TestSummary_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/rating/summary", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"groups": [
								{
									"group_name": "string",
									"items": [
										{
											"change": {
												"direction": "DIRECTION_NONE",
												"meaning": "MEANING_NONE"
											},
											"current_value": 0,
											"name": "string",
											"past_value": 0,
											"rating": "string",
											"rating_direction": "NEUTRAL",
											"status": "OK",
											"value_type": "PERCENT"
										}
									]
								}
							],
							"penalty_score_exceeded": true,
							"premium": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/rating",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Summary(context.Background(), &v1.SummaryRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.SummaryResponse{
		Groups: []v1.SummaryResponseGroup{
			{
				GroupName: "string",
				Items: []v1.SummaryResponseGroupItem{
					{
						Change: v1.SummaryResponseGroupItemChange{
							Direction: v1.SummaryResponseGroupItemChangeDirectionDIRECTIONNONE,
							Meaning:   v1.SummaryResponseGroupItemChangeMeaningMEANINGNONE,
						},
						CurrentValue:    0,
						Name:            "string",
						PastValue:       0,
						Rating:          "string",
						RatingDirection: v1.SummaryResponseGroupItemRatingDirectionNEUTRAL,
						Status:          v1.SummaryResponseGroupItemStatusOK,
						ValueType:       v1.SummaryResponseGroupItemValueTypePERCENT,
					},
				},
			},
		},
		PenaltyScoreExceeded: true,
		Premium:              true,
	}, resp)
}

func TestHistory_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/rating/history", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"date_from":"2019-08-24T14:15:22Z","date_to":"2019-08-24T14:15:22Z","ratings":["string"],"with_premium_scores":true}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"premium_scores": [
								{
									"rating": "string",
									"scores": [
										{
											"date": "2019-08-24T14:15:22Z",
											"rating_value": 0,
											"value": 0
										}
									]
								}
							],
							"ratings": [
								{
									"danger_threshold": 0,
									"premium_threshold": 0,
									"rating": "string",
									"values": [
										{
											"date_from": "2019-08-24T14:15:22Z",
											"date_to": "2019-08-24T14:15:22Z",
											"status": {
												"danger": true,
												"premium": true,
												"warning": true
											},
											"value": 0
										}
									],
									"warning_threshold": 0
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/rating",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.History(context.Background(), &v1.HistoryRequest{
		DateFrom: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
		DateTo:   test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
		Ratings: []string{
			"string",
		},
		WithPremiumScores: true,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.HistoryResponse{
		PremiumScores: []v1.HistoryResponsePremiumScore{
			{
				Rating: "string",
				Scores: []v1.HistoryResponsePremiumScoreScore{
					{
						Date:        test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
						RatingValue: 0,
						Value:       0,
					},
				},
			},
		},
		Ratings: []v1.HistoryResponseRating{
			{
				DangerThreshold:  0,
				PremiumThreshold: 0,
				Rating:           "string",
				Values: []v1.HistoryResponseRatingValue{
					{
						DateFrom: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
						DateTo:   test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
						Status: v1.HistoryResponseRatingValueStatus{
							Danger:  true,
							Premium: true,
							Warning: true,
						},
						Value: 0,
					},
				},
				WarningThreshold: 0,
			},
		},
	}, resp)
}

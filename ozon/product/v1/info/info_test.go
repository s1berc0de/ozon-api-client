package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/diphantxm/ozon-api-client/ozon/product/v1/info"
	"github.com/stretchr/testify/require"
)

func TestDescription_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/info/description", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"offer_id":"5","product_id":73453843}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"id": 73453843,
								"offer_id": "5",
								"name": "Онлайн курс по дрессировке собак \"Воспитанная собака за 3 недели\"",
								"description": "Экспресс-курс - это сокращённый вариант курса \"Собака: инструкция по применению\", дающий базовый минимум знаний, навыков, умений. Это оптимальный вариант для совершения первых шагов по воспитанию!<br/><br/>Что дает Экспресс-курс:<ul><li>Контакт с собакой </li></ul>К концу экспресс-курса дрессировки вы получаете воспитанного друга и соратника, который ориентируется на вас в любой ситуации.<ul><li>Уверенность в безопасности</li></ul>Благополучие собаки: больше не будет срывов с поводка, преследования кошек, попыток съесть что-либо на улице и т. д.<ul><li>Комфортная жизнь</li></ul>Принципиально другой уровень общения, без раздражения, криков и недовольства поведением животного."
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Description(context.Background(), &info.DescriptionRequest{
		OfferID:   "5",
		ProductID: 73453843,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.DescriptionResponse{
		Result: info.DescriptionResponseResult{
			ID:          73453843,
			OfferID:     "5",
			Name:        "Онлайн курс по дрессировке собак \"Воспитанная собака за 3 недели\"",
			Description: "Экспресс-курс - это сокращённый вариант курса \"Собака: инструкция по применению\", дающий базовый минимум знаний, навыков, умений. Это оптимальный вариант для совершения первых шагов по воспитанию!<br/><br/>Что дает Экспресс-курс:<ul><li>Контакт с собакой </li></ul>К концу экспресс-курса дрессировки вы получаете воспитанного друга и соратника, который ориентируется на вас в любой ситуации.<ul><li>Уверенность в безопасности</li></ul>Благополучие собаки: больше не будет срывов с поводка, преследования кошек, попыток съесть что-либо на улице и т. д.<ul><li>Комфортная жизнь</li></ul>Принципиально другой уровень общения, без раздражения, криков и недовольства поведением животного.",
		},
	}, resp)
}

func TestSubscription_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/info/subscription", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"skus":["string"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"count": 0,
									"sku": 0
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Subscription(context.Background(), &info.SubscriptionRequest{
		SKUs: []string{
			"string",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.SubscriptionResponse{
		Result: []info.SubscriptionResponseResult{
			{
				Count: 0,
				SKU:   0,
			},
		},
	}, resp)
}

func TestDiscounted_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/info/discounted", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"discounted_skus":["635548518"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"items": [
								{
									"discounted_sku": 635548518,
									"sku": 320067758,
									"condition_estimation": "4",
									"packaging_violation": "",
									"warranty_type": "",
									"reason_damaged": "Механическое повреждение",
									"comment_reason_damaged": "повреждена заводская упаковка",
									"defects": "",
									"mechanical_damage": "",
									"package_damage": "",
									"shortage": "",
									"repair": "",
									"condition": ""
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Discounted(context.Background(), &info.DiscountedRequest{
		DiscountedSKUs: []string{
			"635548518",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.DiscountedResponse{
		Items: []info.DiscountedResponseItem{
			{
				DiscountedSKU:        635548518,
				SKU:                  320067758,
				ConditionEstimation:  "4",
				PackagingViolation:   "",
				WarrantyType:         "",
				ReasonDamaged:        "Механическое повреждение",
				CommentReasonDamaged: "повреждена заводская упаковка",
				Defects:              "",
				MechanicalDamage:     "",
				PackageDamage:        "",
				Shortage:             "",
				Repair:               "",
				Condition:            "",
			},
		},
	}, resp)
}

package v1_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	v1 "github.com/diphantxm/ozon-api-client/ozon/product/v1"
	"github.com/stretchr/testify/require"
)

func TestImportBySKU_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import-by-sku", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"items":[{"sku":298789742,"name":"string","offer_id":"91132","currency_code":"RUB","old_price":"2590","price":"2300","premium_price":"2200","vat":"0.1"}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  	"result": {
								"task_id": 176594213,
								"unmatched_sku_list": [ ]
						  	}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.ImportBySKU(context.Background(), &v1.ImportBySKURequest{
		Items: []v1.ImportBySKURequestItem{
			{
				SKU:          298789742,
				Name:         "string",
				OfferID:      "91132",
				OldPrice:     "2590",
				Price:        "2300",
				PremiumPrice: "2200",
				Vat:          "0.1",
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.ImportBySKUResponse{
		Result: v1.ImportBySKUResponseResult{
			TaskID:           176594213,
			UnmatchedSKUList: make([]int64, 0),
		},
	}, resp)
}

func TestRatingBySKU_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/rating-by-sku", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"skus":["179737222"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  	"products": [
								{
							  		"sku": 179737222,
							  		"rating": 42.5,
							  		"groups": [
										{
											"key": "media",
								  			"name": "Медиа",
								  			"rating": 70,
								  			"weight": 25,
								  			"conditions": [
												{
									  				"key": "media_images_2",
									  				"description": "Добавлено 2 изображения",
									  				"fulfilled": true,
									  				"cost": 50
												},
												{
													"key": "media_images_3",
													"description": "Добавлено 3 изображения и более",
													"fulfilled": true,
													"cost": 20
												},
												{
									  				"key": "media_image_3d",
									  				"description": "Добавлено 3D-изображение",
									  				"fulfilled": false,
									  				"cost": 15
												},
												{
									  				"key": "media_video",
									  				"description": "Добавлено видео",
									  				"fulfilled": false,
									  				"cost": 15
												}
								  			],
								  			"improve_attributes": [
												{
									  				"id": 4074,
									  				"name": "Код ролика на YouTube"
												},
												{
									  				"id": 4080,
									  				"name": "3D-изображение"
												}
								  			],
								  			"improve_at_least": 2
										},
										{
								  			"key": "important_attributes",
								  			"name": "Важные атрибуты",
								  			"rating": 50,
								  			"weight": 30,
								  			"conditions": [
												{
									  				"key": "important_2",
									  				"description": "Заполнено 2 атрибута и более",
									  				"fulfilled": true,
									  				"cost": 50
												},
												{
									  				"key": "important_50_percent",
									  				"description": "Заполнено более 50% атрибутов",
									  				"fulfilled": false,
									  				"cost": 25
												},
												{
									  				"key": "important_70_percent",
									  				"description": "Заполнено более 70% атрибутов",
									  				"fulfilled": false,
									  				"cost": 25
												}
								  			],
								  			"improve_attributes": [
												{
									  				"id": 4385,
									  				"name": "Гарантийный срок"
												},
												{
									  				"id": 4389,
									  				"name": "Страна-изготовитель"
												},
												{
												  	"id": 8513,
									  				"name": "Количество в упаковке, шт"
												},
												{
											  		"id": 8590,
									  				"name": "Макс. диагональ, дюймы"
												},
												{
												  	"id": 8591,
												  	"name": "Мин. диагональ, дюймы"
												},
												{
												  	"id": 9336,
												  	"name": "Модель браслета/умных часов"
												},
												{
												  	"id": 11046,
												  	"name": "Покрытие"
												},
												{
												  	"id": 11047,
												  	"name": "Прозрачность покрытия"
												},
												{
												  	"id": 11048,
												  	"name": "Дополнительные свойства покрытия"
												},
												{
												  	"id": 11049,
												  	"name": "Вид стекла"
												},
												{
												  	"id": 11603,
												  	"name": "Размер циферблата"
												}
											],
											"improve_at_least": 6
										},
										{
								  			"key": "other_attributes",
								  			"name": "Остальные атрибуты",
											"rating": 0,
											"weight": 25,
											"conditions": [
												{
													"key": "other_2",
													"description": "Заполнено 2 атрибута и более",
													"fulfilled": false,
													"cost": 50
												},
												{
													"key": "other_50_percent",
													"description": "Заполнено более 50% атрибутов",
													"fulfilled": false,
													"cost": 50
												}
											],
								  			"improve_attributes": [
												{
									  				"id": 4382,
									  				"name": "Размеры, мм"
												}
								  			],
								  			"improve_at_least": 1
										},
										{
											"key": "text",
											"name": "Текстовое описание",
											"rating": 50,
											"weight": 20,
											"conditions": [
												{
												  	"key": "text_annotation_100_chars",
												  	"description": "Аннотация более 100 знаков",
												  	"fulfilled": true,
												  	"cost": 25
												},
												{
												  	"key": "text_annotation_500_chars",
												  	"description": "Аннотация более 500 знаков",
												  	"fulfilled": true,
												  	"cost": 25
												},
												{
												  	"key": "text_rich",
												  	"description": "Заполнен Rich-контент",
												  	"fulfilled": false,
												  	"cost": 100
												}
										  	],
											"improve_attributes": [
												{
									  				"id": 11254,
									  				"name": "Rich-контент JSON"
												}
								  			],
								  			"improve_at_least": 1
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
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.RatingBySKU(context.Background(), &v1.RatingBySKURequest{
		SKUs: []string{
			"179737222",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.RatingBySKUResponse{
		Products: []v1.RatingBySKUResponseProduct{
			{
				SKU:    179737222,
				Rating: 42.5,
				Groups: []v1.RatingBySKUResponseProductGroup{
					{
						Key:    "media",
						Name:   "Медиа",
						Rating: 70,
						Weight: 25,
						Conditions: []v1.RatingBySKUResponseProductGroupCondition{
							{
								Key:         "media_images_2",
								Description: "Добавлено 2 изображения",
								Fulfilled:   true,
								Cost:        50,
							},
							{
								Key:         "media_images_3",
								Description: "Добавлено 3 изображения и более",
								Fulfilled:   true,
								Cost:        20,
							},
							{
								Key:         "media_image_3d",
								Description: "Добавлено 3D-изображение",
								Fulfilled:   false,
								Cost:        15,
							},
							{
								Key:         "media_video",
								Description: "Добавлено видео",
								Fulfilled:   false,
								Cost:        15,
							},
						},
						ImproveAttributes: []v1.RatingBySKUResponseProductImproveAttribute{
							{
								ID:   4074,
								Name: "Код ролика на YouTube",
							},
							{
								ID:   4080,
								Name: "3D-изображение",
							},
						},
						ImproveAtLeast: 2,
					},
					{
						Key:    "important_attributes",
						Name:   "Важные атрибуты",
						Rating: 50,
						Weight: 30,
						Conditions: []v1.RatingBySKUResponseProductGroupCondition{
							{
								Key:         "important_2",
								Description: "Заполнено 2 атрибута и более",
								Fulfilled:   true,
								Cost:        50,
							},
							{
								Key:         "important_50_percent",
								Description: "Заполнено более 50% атрибутов",
								Fulfilled:   false,
								Cost:        25,
							},
							{
								Key:         "important_70_percent",
								Description: "Заполнено более 70% атрибутов",
								Fulfilled:   false,
								Cost:        25,
							},
						},
						ImproveAttributes: []v1.RatingBySKUResponseProductImproveAttribute{
							{
								ID:   4385,
								Name: "Гарантийный срок",
							},
							{
								ID:   4389,
								Name: "Страна-изготовитель",
							},
							{
								ID:   8513,
								Name: "Количество в упаковке, шт",
							},
							{
								ID:   8590,
								Name: "Макс. диагональ, дюймы",
							},
							{
								ID:   8591,
								Name: "Мин. диагональ, дюймы",
							},
							{
								ID:   9336,
								Name: "Модель браслета/умных часов",
							},
							{
								ID:   11046,
								Name: "Покрытие",
							},
							{
								ID:   11047,
								Name: "Прозрачность покрытия",
							},
							{
								ID:   11048,
								Name: "Дополнительные свойства покрытия",
							},
							{
								ID:   11049,
								Name: "Вид стекла",
							},
							{
								ID:   11603,
								Name: "Размер циферблата",
							},
						},
						ImproveAtLeast: 6,
					},
					{
						Key:    "other_attributes",
						Name:   "Остальные атрибуты",
						Rating: 0,
						Weight: 25,
						Conditions: []v1.RatingBySKUResponseProductGroupCondition{
							{
								Key:         "other_2",
								Description: "Заполнено 2 атрибута и более",
								Fulfilled:   false,
								Cost:        50,
							},
							{
								Key:         "other_50_percent",
								Description: "Заполнено более 50% атрибутов",
								Fulfilled:   false,
								Cost:        50,
							},
						},
						ImproveAttributes: []v1.RatingBySKUResponseProductImproveAttribute{
							{
								ID:   4382,
								Name: "Размеры, мм",
							},
						},
						ImproveAtLeast: 1,
					},
					{
						Key:    "text",
						Name:   "Текстовое описание",
						Rating: 50,
						Weight: 20,
						Conditions: []v1.RatingBySKUResponseProductGroupCondition{
							{
								Key:         "text_annotation_100_chars",
								Description: "Аннотация более 100 знаков",
								Fulfilled:   true,
								Cost:        25,
							},
							{
								Key:         "text_annotation_500_chars",
								Description: "Аннотация более 500 знаков",
								Fulfilled:   true,
								Cost:        25,
							},
							{
								Key:         "text_rich",
								Description: "Заполнен Rich-контент",
								Fulfilled:   false,
								Cost:        100,
							},
						},
						ImproveAttributes: []v1.RatingBySKUResponseProductImproveAttribute{
							{
								ID:   11254,
								Name: "Rich-контент JSON",
							},
						},
						ImproveAtLeast: 1,
					},
				},
			},
		},
	}, resp)
}

func TestArchive_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/archive", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"product_id":["125529926"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Archive(context.Background(), &v1.ArchiveRequest{
		ProductID: []string{
			"125529926",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.ArchiveResponse{
		Result: true,
	}, resp)
}

func TestUnArchive_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/unarchive", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"product_id":["125529926"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.UnArchive(context.Background(), &v1.UnArchiveRequest{
		ProductID: []string{
			"125529926",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.ArchiveResponse{
		Result: true,
	}, resp)
}

func TestUploadDigitalCodes_Success(t *testing.T) {
	c := v1.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/upload_digital_codes", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"digital_codes":["764282654334"],"product_id":73160317}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"task_id": 172549811
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.UploadDigitalCodes(context.Background(), &v1.UploadDigitalCodesRequest{
		DigitalCodes: []string{
			"764282654334",
		},
		ProductID: 73160317,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v1.UploadDigitalCodesResponse{
		Result: v1.UploadDigitalCodesResponseResult{
			TaskID: 172549811,
		},
	}, resp)
}

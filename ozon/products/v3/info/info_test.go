package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/products/v3/info"
	"github.com/stretchr/testify/require"
)

func TestAttributes_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/products/info/attributes", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"product_id":["213761435"],"visibility":"ALL"},"limit":100,"last_id":"okVsfA==«","sort_dir":"ASC"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"id": 213761435,
									"barcode": "",
									"category_id": 17038062,
									"name": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G",
									"offer_id": "21470",
									"height": 10,
									"depth": 210,
									"width": 140,
									"dimension_unit": "mm",
									"weight": 50,
									"weight_unit": "g",
									"images": [
										{
											"file_name": "https://cdn1.ozone.ru/s3/multimedia-f/6190456071.jpg",
											"default": true,
											"index": 0
										},
										{
											"file_name": "https://cdn1.ozone.ru/s3/multimedia-7/6190456099.jpg",
											"default": false,
											"index": 1
										},
										{
											"file_name": "https://cdn1.ozone.ru/s3/multimedia-9/6190456065.jpg",
											"default": false,
											"index": 2
										}
									],
									"image_group_id": "",
									"images360": [],
									"pdf_list": [],
									"attributes": [
										{
											"attribute_id": 5219,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 970718176,
													"value": "универсальный"
												}
											]
										},
										{
											"attribute_id": 11051,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 970736931,
													"value": "Прозрачный"
												}
											]
										},
										{
											"attribute_id": 10100,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "false"
												}
											]
										},
										{
											"attribute_id": 11794,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 970860783,
													"value": "safe"
												}
											]
										},
										{
											"attribute_id": 9048,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G"
												}
											]
										},
										{
											"attribute_id": 5076,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 39638,
													"value": "Xiaomi"
												}
											]
										},
										{
											"attribute_id": 9024,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "21470"
												}
											]
										},
										{
											"attribute_id": 10015,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "false"
												}
											]
										},
										{
											"attribute_id": 85,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 971034861,
													"value": "Brand"
												}
											]
										},
										{
											"attribute_id": 9461,
											"complex_id": 0,
											"values": [
												{
												"dictionary_value_id": 349824787,
												"value": "Защитная пленка для смартфона"
												}
											]
										},
										{
											"attribute_id": 4180,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G"
												}
											]
										},
										{
											"attribute_id": 4191,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 0,
													"value": "Пленка предназначена для модели Xiaomi Redmi Note 10 Pro 5G. Защитная гидрогелевая пленка обеспечит защиту вашего смартфона от царапин, пыли, сколов и потертостей."
												}
											]
										},
										{
											"attribute_id": 8229,
											"complex_id": 0,
											"values": [
												{
													"dictionary_value_id": 91521,
													"value": "Защитная пленка"
												}
											]
										}
									],
									"complex_attributes": [],
									"color_image": "",
									"last_id": ""
								}
							],
							"total": 1,
							"last_id": "onVsfA=="
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/products/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Attributes(context.Background(), &info.AttributesRequest{
		Filter: info.AttributesRequestFilter{
			ProductID: []string{
				"213761435",
			},
			Visibility: info.AttributesRequestFilterVisibilityALL,
		},
		Limit:   100,
		LastID:  "okVsfA==«",
		SortDir: "ASC",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.AttributesResponse{
		Result: []info.AttributesResponseResult{
			{
				ID:            213761435,
				Barcode:       "",
				CategoryID:    17038062,
				Name:          "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G",
				OfferID:       "21470",
				Height:        10,
				Depth:         210,
				Width:         140,
				DimensionUnit: info.AttributesResponseResultDimensionUnitMm,
				Weight:        50,
				WeightUnit:    "g",
				Images: []info.AttributesResponseResultImage{
					{
						FileName: "https://cdn1.ozone.ru/s3/multimedia-f/6190456071.jpg",
						Default:  true,
						Index:    0,
					},
					{
						FileName: "https://cdn1.ozone.ru/s3/multimedia-7/6190456099.jpg",
						Default:  false,
						Index:    1,
					},
					{
						FileName: "https://cdn1.ozone.ru/s3/multimedia-9/6190456065.jpg",
						Default:  false,
						Index:    2,
					},
				},
				ImageGroupID: "",
				Images360:    make([]info.AttributesResponseResultImage360, 0),
				PDFList:      make([]info.AttributesResponseResultPDFItem, 0),
				Attributes: []info.AttributesResponseResultAttribute{
					{
						AttributeID: 5219,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 970718176,
								Value:             "универсальный",
							},
						},
					},
					{
						AttributeID: 11051,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 970736931,
								Value:             "Прозрачный",
							},
						},
					},
					{
						AttributeID: 10100,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "false",
							},
						},
					},
					{
						AttributeID: 11794,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 970860783,
								Value:             "safe",
							},
						},
					},
					{
						AttributeID: 9048,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G",
							},
						},
					},
					{
						AttributeID: 5076,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 39638,
								Value:             "Xiaomi",
							},
						},
					},
					{
						AttributeID: 9024,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "21470",
							},
						},
					},
					{
						AttributeID: 10015,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "false",
							},
						},
					},
					{
						AttributeID: 85,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 971034861,
								Value:             "Brand",
							},
						},
					},
					{
						AttributeID: 9461,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 349824787,
								Value:             "Защитная пленка для смартфона",
							},
						},
					},
					{
						AttributeID: 4180,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G",
							},
						},
					},
					{
						AttributeID: 4191,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 0,
								Value:             "Пленка предназначена для модели Xiaomi Redmi Note 10 Pro 5G. Защитная гидрогелевая пленка обеспечит защиту вашего смартфона от царапин, пыли, сколов и потертостей.",
							},
						},
					},
					{
						AttributeID: 8229,
						ComplexID:   0,
						Values: []info.AttributesResponseResultAttributeValue{
							{
								DictionaryValueID: 91521,
								Value:             "Защитная пленка",
							},
						},
					},
				},
				ComplexAttributes: make([]info.AttributesResponseResultComplexAttribute, 0),
				ColorImage:        "",
			},
		},
		Total:  1,
		LastID: "onVsfA==",
	}, resp)
}

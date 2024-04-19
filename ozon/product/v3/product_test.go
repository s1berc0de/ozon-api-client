package v3_test

import (
	"bytes"
	"context"
	"github.com/s1berc0de/ozon-api-client/internal/auth"
	"github.com/s1berc0de/ozon-api-client/internal/test"
	v3 "github.com/s1berc0de/ozon-api-client/ozon/product/v3"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestImport_Success(t *testing.T) {
	c := v3.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/product/import", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"items":[{"attributes":[{"complex_id":0,"id":5076,"values":[{"dictionary_value_id":971082156,"value":"Стойка для акустической системы"}]},{"complex_id":0,"id":10096,"values":[{"dictionary_value_id":61576,"value":"серый"}]}],"barcode":"112772873170","description_category_id":17033876,"color_image":"","complex_attributes":[],"currency_code":"RUB","depth":10,"dimension_unit":"nm","height":250,"images":[],"images360":[],"name":"Комплект защитных плёнок для X3 NFC. Темный хлопок","offer_id":"143210608","old_price":"1100","pdf_list":[],"premium_price":"900","price":"1000","primary_image":"","vat":"0.1","weight":100,"weight_unit":"g","width":150}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(bytes.NewBufferString(`{"result":{"task_id":172549793}}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Import(context.Background(), &v3.ImportRequest{
		Items: []v3.ImportItem{
			{
				Attributes: []v3.ImportItemAttribute{
					{
						ComplexID: 0,
						ID:        5076,
						Values: []v3.ImportItemAttributeValue{
							{
								DictionaryValueID: 971082156,
								Value:             "Стойка для акустической системы",
							},
						},
					},
					{
						ComplexID: 0,
						ID:        10096,
						Values: []v3.ImportItemAttributeValue{
							{
								DictionaryValueID: 61576,
								Value:             "серый",
							},
						},
					},
				},
				Barcode:               "112772873170",
				DescriptionCategoryID: 17033876,
				ColorImage:            "",
				ComplexAttributes:     []v3.ComplexAttribute{},
				CurrencyCode:          "RUB",
				Depth:                 10,
				DimensionUnit:         "nm",
				Height:                250,
				Images:                []string{},
				Images360:             []string{},
				Name:                  "Комплект защитных плёнок для X3 NFC. Темный хлопок",
				OfferID:               "143210608",
				OldPrice:              "1100",
				PdfList:               []string{},
				PremiumPrice:          "900",
				Price:                 "1000",
				PrimaryImage:          "",
				Vat:                   "0.1",
				Weight:                100,
				WeightUnit:            "g",
				Width:                 150,
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v3.ImportResponse{
		Result: v3.ImportResult{
			TaskID: 172549793,
		},
	}, resp)
}

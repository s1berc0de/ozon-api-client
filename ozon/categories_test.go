package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetProductTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductTreeParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductTreeParams{
				Language: English,
			},
			`{
				"result": [
				  {
					"description_category_id": 0,
					"category_name": "string",
					"children": [],
					"disabled": true,
					"type_id": 0,
					"type_name": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductTreeParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().Tree(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductTreeResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetCategoryAttributes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetCategoryAttributesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetCategoryAttributesParams{
				DescriptionCategoryId: 12345,
				Language:              English,
				TypeId:                2,
			},
			`{
				"result": [
				  {
					"category_dependent": true,
					"description": "string",
					"dictionary_id": 0,
					"group_id": 0,
					"group_name": "string",
					"id": 0,
					"is_aspect": true,
					"is_collection": true,
					"is_required": true,
					"name": "string",
					"type": "string",
					"attribute_complex_id": 0,
					"max_value_count": 0
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetCategoryAttributesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().Attributes(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetCategoryAttributesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetAttributeDictionary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetAttributeDictionaryParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetAttributeDictionaryParams{
				AttributeId:           123456,
				DescriptionCategoryId: 12,
				Language:              English,
				LastValueId:           1,
				Limit:                 5,
				TypeId:                6,
			},
			`{
				"has_next": true,
				"result": [
				  {
					"id": 0,
					"info": "string",
					"picture": "string",
					"value": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetAttributeDictionaryParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().AttributesDictionary(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetAttributeDictionaryResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > int(test.params.Limit) {
				t.Errorf("Length of response result is bigger than limit")
			}
		}
	}
}

func TestSearchAttributeDictionary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SearchAttributeDictionaryParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SearchAttributeDictionaryParams{
				AttributeId:           123456,
				DescriptionCategoryId: 12,
				Value:                 "34",
				Limit:                 5,
				TypeId:                6,
			},
			`{
				"has_next": true,
				"result": [
				  {
					"id": 0,
					"info": "string",
					"picture": "string",
					"value": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SearchAttributeDictionaryParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().SearchAttributesDictionary(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetAttributeDictionaryResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > int(test.params.Limit) {
				t.Errorf("Length of response result is bigger than limit")
			}
		}
	}
}

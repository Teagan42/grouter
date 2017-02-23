package validators

import (
	"testing"
)

func TestSwaggerDoc(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = SwaggerDoc().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid swagger doc")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []obj{
				obj{
					"pid": obj{},
				},
				obj{
					"/:id": obj{
						"nothing": obj{},
					},
					"\\": obj{
						"name": "apiKey",
						"type": "basic",
						"in": "query",
					},
				},
			}
			for _, value := range testValues {
				func (val obj) {
					var path, err = SwaggerDoc().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid swagger doc - %s", val, path)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []obj{
			obj{
				"swagger": "2.0",
				"info": obj{
					"title": "Test",
					"description": "Testing stuff",
					"contact": obj{"name": "Bob", "url": "google.com", "email": "teagan@massroots.com"},
					"license": obj{"name": "Bob", "url": "google.com"},
				},
				"host": "api.massroots.com",
				"basePath": "/v1",
				"schemes": []string{"http", "https"},
				"consumes": []string{"application/json"},
				"produces": []string{"application/json"},
				"paths": obj{
					"/": obj{
						"get": obj{
							"summary": "title",
							"description": "desc",
							"deprecated": false,
						},
					},
					"/posts/:id/": obj{
						"get": obj{
							"summary": "get post",
							"description": "get post",
							"deprecated": true,
						},
						"post": obj{
							"summary": "create post",
							"description": "create post",
							"deprecated": true,
						},
					},
				},
				"securityDefinitions":obj{
					"apiKey": obj{
						"name": "api",
						"in": "header",
						"type": "basic",
					},
					"admin": obj{
						"name": "isAdmin",
						"in": "query",
						"type": "basic",
					},
				},
			},
		}
		for _, value := range testValues {
			func (val obj) {
				var _, err = SwaggerDoc().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid swagger doc", val);
				}
			} (value)
		}
	})
}

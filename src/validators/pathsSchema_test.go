package validators

import (
	"testing"
)

func TestPaths(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = Paths().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid paths")
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
						"get": obj{
							"summary": "test",
							"description": "test",
							"deprecated": false,
						},
					},
				},
			}
			for _, value := range testValues {
				func (val obj) {
					var _, err = HttpMethod().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid http method", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []obj{
			obj{
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
		}
		for _, value := range testValues {
			func (val obj) {
				var _, err = Paths().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid paths", val);
				}
			} (value)
		}
	})
}

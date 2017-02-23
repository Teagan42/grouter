package validators

import (
	"testing"
)

func TestPaths(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			_, err := Paths().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid paths - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
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
				func (val *interface{}) {
					_, err := Paths().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid paths - %s", val, err)
					}
				} (&value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		testValues := []interface{}{
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
			func (val interface{}) {
				_, err := Paths().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid paths - %s", val, err);
				}
			} (value)
		}
	})
}

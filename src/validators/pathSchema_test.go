package validators

import (
	"testing"
)

func TestHttpMethod(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = HttpMethod().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid paths")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []string{"", "test", "@gmail", "@.com", "@gmail.com", "application/"}
			for _, value := range testValues {
				func (val string) {
					var _, err = HttpMethod().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an paths", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []string{"get", "put", "post", "delete"}
		for _, value := range testValues {
			func (val string) {
				var _, err = HttpMethod().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid paths", val);
				}
			} (value)
		}
	})
}

func TestPath(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = Path().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid path")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []obj{
				obj{
					"get": obj{
						"description": "Test",
						"deprecated": true,
					},
					"post": 5,
				},
			}
			for _, value := range testValues {
				func (val obj) {
					var _, err = Path().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid path", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []obj{
			obj{
				"delete": obj{
					"summary": "test",
					"description": "Testing",
					"deprecated": true,
				},
			},
		}
		for _, value := range testValues {
			func (val obj) {
				var _, err = Path().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid path", val);
				}
			} (value)
		}
	})
}

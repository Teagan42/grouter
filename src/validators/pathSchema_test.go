package validators

import (
	"testing"
)

func TestHttpMethod(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var res, err = HttpMethod().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid paths - %s", res)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{"", "test", "@gmail", "@.com", "@gmail.com", "application/"}
			for _, value := range testValues {
				func (val interface{}) {
					var res, err = HttpMethod().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an paths - %s", val, res)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []interface{}{"get", "put", "post", "delete"}
		for _, value := range testValues {
			func (val interface{}) {
				var res, err = HttpMethod().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid paths - %s", val, res);
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
			_, err := Path().Validate(nil)

			if err == nil {
				t.Errorf("Expected nil to be an invalid path - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
				obj{
					"get": obj{
						"description": "Test",
						"deprecated": true,
					},
					"post": 5,
				},
			}
			for _, value := range testValues {
				func (val interface{}) {
					_, err := Path().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid path - %s", val, err)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		testValues := []interface{}{
			obj{
				"delete": obj{
					"summary": "test",
					"description": "Testing",
					"deprecated": true,
				},
			},
		}
		for _, value := range testValues {
			func (val interface{}) {
				_, err := Path().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid path - %s", val, err);
				}
			} (value)
		}
	})
}

package validators

import (
	"testing"
	v "github.com/gima/govalid/v1"
)

func TestOperation(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = Operation().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid operation")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []obj{
				obj{},
				obj{"tags": []string{}},
			}
			for _, value := range testValues {
				func (val obj) {
					var _, err = Operation().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid operation", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []obj{
			obj{
				"summary": "test",
				"description": "Testing",
				"deprecated": false,
			},
		}
		for _, value := range testValues {
			func (val obj) {
				var _, err = Operation().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid mime type", val);
				}
			} (value)
		}
	})
}

func TestStringMimeType(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = v.String(StringMimeType()).Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid mime type")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []string{"", "test", "@gmail", "@.com", "@gmail.com", "application/"}
			for _, value := range testValues {
				func (val string) {
					var _, err = v.String(StringMimeType()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid mimetype", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []string{"application/*", "image/jpeg", "application/json"}
		for _, value := range testValues {
			func (val string) {
				var _, err = v.String(StringMimeType()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid mime type", val);
				}
			} (value)
		}
	})
}
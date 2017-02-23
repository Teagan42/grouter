package validators

import (
"testing"
v "github.com/gima/govalid/v1"
)

func TestStringUrl(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = v.String(StringUrl()).Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid url")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{"", "%google", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val interface{}) {
					var _, err = v.String(StringUrl()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid url - %s", val, err)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []interface{}{"google.com", "massroots.com", "http://teagantotally.rocks"}

		for _, value := range testValues {
			func (val interface{}) {
				var _, err = v.String(StringUrl()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid url - %s", val, err);
				}
			} (value)
		}
	})
}

func TestStringRelative(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = v.String(StringRelative()).Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid relative url - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{"", "test", "%google", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val interface{}) {
					var _, err = v.String(StringRelative()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid relative url - %s", val, err)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []interface{}{"/google.com", "/posts/", "/api/comments/203"}

		for _, value := range testValues {
			func (val interface{}) {
				var _, err = v.String(StringRelative()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid relative url - %s", val, err);
				}
			} (value)
		}
	})
}
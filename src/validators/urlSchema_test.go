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
			var testValues = []string{"", "%google", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val string) {
					var _, err = v.String(StringUrl()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid url", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []string{"google.com", "massroots.com", "http://teagantotally.rocks"}
		for _, value := range testValues {
			func (val string) {
				var _, err = v.String(StringUrl()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid url", val);
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
				t.Error("Expected nil to be an invalid relative url")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []string{"", "test", "%google", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val string) {
					var _, err = v.String(StringRelative()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid relative url", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []string{"/google.com", "/posts/", "/api/comments/203"}
		for _, value := range testValues {
			func (val string) {
				var _, err = v.String(StringRelative()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid relative url", val);
				}
			} (value)
		}
	})
}
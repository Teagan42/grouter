package validators

import (
	"testing"
	v "github.com/gima/govalid/v1"
)

func TestStringEmail(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = v.String(StringEmail()).Validate(nil)

			if err == nil {
				t.Errorf("Expected nil to be an invalid email - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{"", "test", "@gmail", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val interface{}) {
					var _, err = v.String(StringEmail()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid email - %s", val, err)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []interface{}{"teagan@massroots.com", "t.m.g@me.co", "that@teagantotally.rocks"}

		for _, value := range testValues {
			func (val interface{}) {
				var _, err = v.String(StringEmail()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid email - %s", val, err);
				}
			} (value)
		}
	})
}
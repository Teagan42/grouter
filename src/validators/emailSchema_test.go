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
				t.Error("Expected nil to be an invalid email")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []string{"", "test", "@gmail", "@.com", "@gmail.com"}
			for _, value := range testValues {
				func (val string) {
					var _, err = v.String(StringEmail()).Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid email", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []string{"teagan@massroots.com", "t.m.g@me.co", "that@teagantotally.rocks"}
		for _, value := range testValues {
			func (val string) {
				var _, err = v.String(StringEmail()).Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid email", val);
				}
			} (value)
		}
	})
}
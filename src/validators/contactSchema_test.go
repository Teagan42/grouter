package validators

import (
	"testing"
	"encoding/json"
)

func TestContactSchema(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()

			_, err := Contact().Validate(nil)

			if err == nil {
				t.Errorf("expected nil to be an invalid contact: %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
				obj{"url": "google.com", "email": "teagan@massroots.com"},
				obj{"name": 5, "email": 0.5, "url": "google.com"},
				obj{"name": "Bob", "email": "nope", "url": "here"},
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := Contact().Validate(val)

					if err == nil {
						disp, _ := json.Marshal(val)
						t.Errorf("expecteded %s to be an invalid contact - %s", disp, err)
					}
				}(value)
			}
		})
	})
	t.Run("validates", func (t *testing.T) {
		t.Parallel()

		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
				obj{"name": "Bob", "url": "google.com", "email": "teagan@massroots.com"},
				obj{"name": "Jim", "email": "123eye@on.me", "url": "https://massroots.com"},
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := Contact().Validate(val)

					if err != nil {
						disp, _ := json.Marshal(val)
						t.Errorf("expecteded %s to be a valid contact - %s", disp, err)
					}
				}(value)
			}
		})
	})
}

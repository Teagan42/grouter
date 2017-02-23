package validators

import (
	"testing"
	"encoding/json"
)

func TestLicenseSchema(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()

			_, err := License().Validate(nil)

			if err == nil {
				t.Error("expected nil to be an invalid license - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
				obj{"url": "google.com"},
				obj{"name": 5, "url": "google.com"},
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := License().Validate(val)

					if err == nil {
						disp, _ := json.Marshal(val)
						t.Errorf("expecteded %s to be an invalid license - %s", disp, err)
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
				obj{"name": "Bob", "url": "google.com"},
				obj{"name": "Jim", "url": "https://massroots.com"},
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := License().Validate(val)

					if err != nil {
						disp, _ := json.Marshal(val)
						t.Errorf("expecteded %s to be a valid license - %s", disp, err)
					}
				}(value)
			}
		})
	})
}

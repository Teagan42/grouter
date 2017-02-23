package validators

import (
	"testing"
	"validators"
	"encoding/json"
)

func TestInfo(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()

			_, err := validators.Info().Validate(nil)

			if err == nil {
				t.Error("expected nil to be an invalid license - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			testValues := []interface{}{
				obj{"url": "google.com"},
				obj{"name": 5, "url": "google.com"},
				//obj{"name": "Bob", "url": "here"}, TODO : Why is this not invalid?
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := validators.Info().Validate(val)

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
				obj{
					"title": "Test",
					"description": "Testing stuff",
					"contact": obj{"name": "Bob", "url": "google.com", "email": "teagan@massroots.com"},
					"license": obj{"name": "Bob", "url": "google.com"},
				},
			}

			for _, value := range testValues {
				func (val interface{}) {
					_, err := validators.Info().Validate(val)

					if err != nil {
						disp, _ := json.Marshal(val)
						t.Errorf("expecteded %s to be a valid license - %s", disp, err)
					}
				}(value)
			}
		})
	})
}

package validators

import (
	"testing"
)

func TestSecurityDefinition(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()
			var _, err = SecurityDefinition().Validate(nil)

			if err == nil {
				t.Error("Expected nil to be an invalid security definition - %s", err)
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{
				obj{
					"pid": obj{},
				},
				obj{
					"/:id": obj{
						"nothing": obj{},
					},
					"\\": obj{
						"name": "apiKey",
						"type": "basic",
						"in": "query",
					},
				},
			}
			for _, value := range testValues {
				func (val interface{}) {
					var _, err = SecurityDefinition().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid security definition", val, err)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []interface{}{
			obj{
				"apiKey": obj{
					"name": "api",
					"in": "header",
					"type": "basic",
				},
				"admin": obj{
					"name": "isAdmin",
					"in": "query",
					"type": "basic",
				},
			},
		}
		for _, value := range testValues {
			func (val interface{}) {
				var _, err = SecurityDefinition().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid security definition - %s", val, err);
				}
			} (value)
		}
	})
}

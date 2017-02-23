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
				t.Error("Expected nil to be an invalid security definition")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []obj{
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
				func (val obj) {
					var _, err = SecurityDefinition().Validate(val)

					if err == nil {
						t.Errorf("Expected '%s' to be an invalid security definition", val)
					}
				} (value)
			}
		})
	})
	t.Run("validates", func(t *testing.T) {
		t.Parallel()
		var testValues = []obj{
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
			func (val obj) {
				var _, err = SecurityDefinition().Validate(val)

				if err != nil {
					t.Errorf("Expected %s to be a valid security definition", val);
				}
			} (value)
		}
	})
}

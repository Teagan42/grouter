package validators

import (
	v "github.com/gima/govalid/v1"
	"testing"
	"encoding/json"
)

func TestParameterType(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()

			var _, err = ParameterType().Validate(nil)

			if err == nil {
				t.Error("expected nil to be an invalid parameter type")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{
				obj{"url": "google.com"},
				obj{"name": 5, "url": "google.com"},
				obj{},
			}

			for _, value := range testValues {
				func (val interface{}) {
					var _, err = ParameterType().Validate(val)

					if err == nil {
						var disp, err = json.Marshal(val)
						t.Errorf("expecteded %s to be an invalid parameter type %s", disp, err)
					}
				}(value)
			}
		})
	})
	t.Run("validates", func (t *testing.T) {
		t.Parallel()

		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{
				v.String(),
				v.Object(),
			}

			for _, value := range testValues {
				func (val interface{}) {
					var r, err = ParameterType().Validate(val)

					if err != nil {
						var disp, _ = json.Marshal(val)
						t.Errorf("expecteded %s to be a valid parameter type %s %s", disp, r, err)
					}
				}(value)
			}
		})
	})
}

func TestParameter(t *testing.T) {
	t.Run("invalidates", func (t *testing.T) {
		t.Parallel()
		t.Run("nil", func (t *testing.T) {
			t.Parallel()

			var _, err = Parameter().Validate(nil)

			if err == nil {
				t.Error("expected nil to be an invalid parameter")
			}
		})
		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{
				obj{"url": "google.com"},
				obj{"name": 5, "url": "google.com"},
				obj{},
			}

			for _, value := range testValues {
				func (val interface{}) {
					var _, err = Parameter().Validate(val)

					if err == nil {
						var disp, err = json.Marshal(val)
						t.Errorf("expecteded %s to be an invalid parameter %s", disp, err)
					}
				}(value)
			}
		})
	})
	t.Run("validates", func (t *testing.T) {
		t.Parallel()

		t.Run("values", func (t *testing.T) {
			t.Parallel()
			var testValues = []interface{}{
				obj{
					"name": "id",
					"in": "path",
					"type": obj{
						"name": "number",
						"validator": v.Number(),
					},
					"description": "Field id",
					"required": true,
				},
			}

			for _, value := range testValues {
				func (val interface{}) {
					var r, err = Parameter().Validate(val)

					if err != nil {
						var disp, _ = json.Marshal(val)
						t.Errorf("expecteded %s to be a valid parameter %s - %s", disp, r, err)
					}
				}(value)
			}
		})
	})
}
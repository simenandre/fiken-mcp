package fiken

import (
	"encoding/json"
	"testing"
)

func TestConvertMoneyFieldsFromOre(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, result map[string]interface{})
	}{
		{
			name:  "converts net, gross, vat fields",
			input: `{"net":10000,"gross":12500,"vat":2500}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "net", 100.0)
				assertFloat(t, result, "gross", 125.0)
				assertFloat(t, result, "vat", 25.0)
			},
		},
		{
			name:  "converts amount field",
			input: `{"amount":49900}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "amount", 499.0)
			},
		},
		{
			name:  "converts unitPrice field",
			input: `{"unitPrice":99900}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "unitPrice", 999.0)
			},
		},
		{
			name:  "converts fields ending in Amount",
			input: `{"netAmount":10000,"grossAmount":12500,"vatAmount":2500}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "netAmount", 100.0)
				assertFloat(t, result, "grossAmount", 125.0)
				assertFloat(t, result, "vatAmount", 25.0)
			},
		},
		{
			name:  "does not convert non-money fields",
			input: `{"quantity":5,"page":0,"customerId":123}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "quantity", 5.0)
				assertFloat(t, result, "page", 0.0)
				assertFloat(t, result, "customerId", 123.0)
			},
		},
		{
			name:  "converts nested money fields",
			input: `{"lines":[{"net":5000,"description":"item"}]}`,
			check: func(t *testing.T, result map[string]interface{}) {
				lines := result["lines"].([]interface{})
				line := lines[0].(map[string]interface{})
				assertFloat(t, line, "net", 50.0)
				if line["description"] != "item" {
					t.Errorf("expected description 'item', got %v", line["description"])
				}
			},
		},
		{
			name:  "returns original on invalid JSON",
			input: `not json`,
			check: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertMoneyFieldsFromOre([]byte(tt.input))
			if tt.check == nil {
				if string(result) != tt.input {
					t.Errorf("expected original input returned unchanged, got %s", result)
				}
				return
			}
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("result is not valid JSON: %v — got: %s", err, result)
			}
			tt.check(t, parsed)
		})
	}
}

func TestConvertMoneyFieldsToOre(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, result map[string]interface{})
	}{
		{
			name:  "converts net, gross, vat fields",
			input: `{"net":100,"gross":125,"vat":25}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "net", 10000)
				assertFloat(t, result, "gross", 12500)
				assertFloat(t, result, "vat", 2500)
			},
		},
		{
			name:  "rounds fractional NOK to nearest øre",
			input: `{"amount":99.999}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "amount", 10000)
			},
		},
		{
			name:  "converts fields ending in Amount",
			input: `{"netAmount":100,"grossAmount":125}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "netAmount", 10000)
				assertFloat(t, result, "grossAmount", 12500)
			},
		},
		{
			name:  "does not convert non-money fields",
			input: `{"quantity":5,"page":0}`,
			check: func(t *testing.T, result map[string]interface{}) {
				assertFloat(t, result, "quantity", 5.0)
				assertFloat(t, result, "page", 0.0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertMoneyFieldsToOre([]byte(tt.input))
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("result is not valid JSON: %v — got: %s", err, result)
			}
			tt.check(t, parsed)
		})
	}
}

func assertFloat(t *testing.T, m map[string]interface{}, key string, expected float64) {
	t.Helper()
	val, ok := m[key]
	if !ok {
		t.Errorf("key %q not found in result", key)
		return
	}
	got, ok := val.(float64)
	if !ok {
		t.Errorf("key %q: expected float64, got %T (%v)", key, val, val)
		return
	}
	if got != expected {
		t.Errorf("key %q: expected %v, got %v", key, expected, got)
	}
}

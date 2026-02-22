package fiken

import (
	"encoding/json"
	"math"
	"strings"
)

// moneyFields contains field names that represent monetary values in øre (smallest currency unit).
var moneyFields = map[string]bool{
	"net":         true,
	"gross":       true,
	"vat":         true,
	"amount":      true,
	"unitPrice":   true,
	"balance":     true,
	"paid":        true,
	"outstanding": true,
	"netInNok":    true,
	"grossInNok":  true,
	"vatInNok":    true,
}

// isMoneyField returns true if the field name represents a monetary value stored in øre.
func isMoneyField(name string) bool {
	if moneyFields[name] {
		return true
	}
	return strings.HasSuffix(name, "Amount")
}

// ConvertMoneyFieldsFromOre converts monetary fields in a JSON payload from øre to NOK
// (divides integer øre values by 100 to produce decimal NOK values).
// Non-JSON input is returned unchanged.
func ConvertMoneyFieldsFromOre(data []byte) []byte {
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return data
	}
	converted, err := json.Marshal(convertValueFromOre(parsed, ""))
	if err != nil {
		return data
	}
	return converted
}

// ConvertMoneyFieldsToOre converts monetary fields in a JSON payload from NOK to øre
// (multiplies decimal NOK values by 100 and rounds to the nearest integer).
// Non-JSON input is returned unchanged.
func ConvertMoneyFieldsToOre(data []byte) []byte {
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return data
	}
	converted, err := json.Marshal(convertValueToOre(parsed, ""))
	if err != nil {
		return data
	}
	return converted
}

func convertValueFromOre(v interface{}, fieldName string) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(val))
		for k, mv := range val {
			result[k] = convertValueFromOre(mv, k)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(val))
		for i, item := range val {
			result[i] = convertValueFromOre(item, fieldName)
		}
		return result
	case float64:
		if isMoneyField(fieldName) {
			return val / 100.0
		}
		return val
	default:
		return val
	}
}

func convertValueToOre(v interface{}, fieldName string) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(val))
		for k, mv := range val {
			result[k] = convertValueToOre(mv, k)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(val))
		for i, item := range val {
			result[i] = convertValueToOre(item, fieldName)
		}
		return result
	case float64:
		if isMoneyField(fieldName) {
			// int64 is marshaled as a JSON integer (no decimal point), which is the
			// format the Fiken API expects for øre values.
			return int64(math.Round(val * 100))
		}
		return val
	default:
		return val
	}
}

package utils

import (
	"encoding/json"
)

func Jsonify(a interface{}) string {
	output, _ := json.MarshalIndent(a, "", "  ")
	return string(output)
}

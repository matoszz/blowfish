package utils

import (
	"time"

	"github.com/tidwall/gjson"
)

func formatJSONValue(v gjson.Result) string {
	if !v.Time().IsZero() {
		// if the value can be formatted into time, assume it's a time
		return v.Time().Local().Format(time.RFC822)
	}

	return v.String()
}

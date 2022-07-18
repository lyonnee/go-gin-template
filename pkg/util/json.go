package util

import (
	"bytes"
	"encoding/json"
)

func CompressJson(jsonStr string) (string, error) {
	dst := bytes.Buffer{}
	if err := json.Compact(&dst, []byte(jsonStr)); err != nil {
		return "", err
	}

	return dst.String(), nil
}

package model

import (
	"encoding/json"
	"log"
)

// JSONResult json result
type JSONResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

// MustUnmarshalResult unmashal result
func MustUnmarshalResult(value []byte) *JSONResult {
	res := &JSONResult{}
	err := json.Unmarshal(value, res)
	if err != nil {
		log.Fatalf("unmashal result failed with %+v", err)
	}

	return res
}

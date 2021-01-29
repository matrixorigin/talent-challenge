package model

import (
	"encoding/json"
)

// Request request
type Request struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

// MustMarshal marshal
func (r *Request) MustMarshal() []byte {
	v, _ := json.Marshal(r)
	return v
}

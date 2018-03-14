package rqf

import (
	"errors"

	"github.com/softbrewery/gojoi/pkg/joi"
)

// Parse error definitions
var (
	ErrBadFormat = errors.New("Bad filter format")
)

// Parser ...
type Parser struct {
}

// Parse parses a raw json filter into Filter object
func Parse(rawFilter string, schema joi.Schema) (*Filter, error) {
	_ = schema
	filter := NewFilter()
	return filter, nil
}

// MustParse parses a raw json and panics on error
func MustParse(rawFilter string, schema joi.Schema) *Filter {
	filter, err := Parse(rawFilter, schema)
	if err != nil {
		panic(err)
	}
	return filter
}

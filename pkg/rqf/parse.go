package rqf

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/softbrewery/gojoi/pkg/joi"
)

// Parse error definitions
var (
	ErrBadFormat = errors.New("Bad filter format")
)

// Parser ...
type Parser struct {
	fieldsSchema *joi.Schema
	orderSchema  *joi.Schema
	limitSchema  *joi.Schema
	offsetSchema *joi.Schema
	whereSchema  *joi.Schema
}

// NewParser ...
func NewParser() *Parser {
	return &Parser{}
}

// FieldsSchema ...
func (p *Parser) FieldsSchema(fieldsSchema joi.Schema) *Parser {
	p.fieldsSchema = &fieldsSchema
	return p
}

// Parse parses a raw json filter into Filter object
func (p *Parser) Parse(rawFilter string) (*Filter, error) {
	normalizedFilter, err := normalizeFilter(rawFilter)
	if err != nil {
		return nil, err
	}

	filter, err := parseJSONFilter(normalizedFilter)
	if err != nil {
		return nil, err
	}

	if joi.IsSet(p.fieldsSchema) {
		err := joi.Validate(filter.Fields, joi.Array().Items(*p.fieldsSchema))
		if err != nil {
			return nil, err
		}
	}

	return filter, nil
}

// MustParse parses a raw json and panics on error
func (p *Parser) MustParse(rawFilter string) *Filter {
	filter, err := p.Parse(rawFilter)
	if err != nil {
		panic(err)
	}
	return filter
}

// normalizeFilter will filter the 'filter' query parameter from the string
func normalizeFilter(rawFilter string) (string, error) {
	stripped := rawFilter

	index := strings.Index(stripped, "?filter=")
	if index != -1 {
		stripped = stripped[index+8:]
	}

	decoded, err := url.QueryUnescape(stripped)

	if err != nil {
		return "", err
	}

	return decoded, nil
}

// parseJSONFilter will parse the json into a Query object
func parseJSONFilter(rawFilter string) (*Filter, error) {
	filter := NewFilter()

	err := json.Unmarshal([]byte(rawFilter), filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

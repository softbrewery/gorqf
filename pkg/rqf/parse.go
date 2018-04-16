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
	fieldSchema  *joi.Schema
	orderSchema  *joi.Schema
	limitSchema  *joi.Schema
	offsetSchema *joi.Schema
	whereSchema  *joi.Schema
}

// NewParser ...
func NewParser() *Parser {
	return &Parser{}
}

// FieldSchema ...
func (p *Parser) FieldSchema(fieldSchema joi.Schema) *Parser {
	p.fieldSchema = &fieldSchema
	return p
}

// OrderSchema ...
func (p *Parser) OrderSchema(orderSchema joi.Schema) *Parser {
	p.orderSchema = &orderSchema
	return p
}

// LimitSchema ...
func (p *Parser) LimitSchema(limitSchema joi.Schema) *Parser {
	p.limitSchema = &limitSchema
	return p
}

// OffsetSchema ...
func (p *Parser) OffsetSchema(offsetSchema joi.Schema) *Parser {
	p.offsetSchema = &offsetSchema
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

	if joi.IsSet(p.fieldSchema) {
		err := joi.Validate(filter.Fields, joi.Slice().Items(*p.fieldSchema))
		if err != nil {
			return nil, err
		}
	}

	if joi.IsSet(p.orderSchema) {
		err := joi.Validate(filter.Order, joi.Slice().Items(*p.orderSchema))
		if err != nil {
			return nil, err
		}
	}

	if joi.IsSet(p.limitSchema) {
		err := joi.Validate(filter.Limit, *p.limitSchema)
		if err != nil {
			return nil, err
		}
	}

	if joi.IsSet(p.offsetSchema) {
		err := joi.Validate(filter.Offset, *p.offsetSchema)
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

	index := strings.Index(stripped, "filter=")
	if index != -1 {
		stripped = stripped[index+7:]
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

package rqf

// Filter keeps info about the filter
type Filter struct {
	Fields []string               `json:"fields"`
	Order  []string               `json:"order"`
	Limit  int                    `json:"limit"`
	Offset int                    `json:"offset"`
	Where  map[string]interface{} `json:"where"`
}

// NewFilter creates a new filter object
func NewFilter() *Filter {
	return &Filter{}
}

// IsEmpty returns true if non of the filter properties are set
func (f *Filter) IsEmpty() bool {
	if len(f.Fields) > 0 ||
		len(f.Order) > 0 ||
		f.Limit > 0 ||
		f.Offset > 0 ||
		f.Where != nil {
		return false
	}
	return true
}

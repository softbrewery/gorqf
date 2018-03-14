package rqf

// Filter keeps info about the filter
type Filter struct {
	Fields []string
	Order  []string
	Limit  int
	Offset int
	Where  interface{}
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

package types

type ListOptions struct {
	// for pagination
	Limit int
	Skip  int

	Filter map[string]bool
}

type ListOption interface {
	apply(*ListOptions)
}

// funcListOption wraps a function that modifies ListOptions into an
// implementation of the ListOption interface.
type funcListOption struct {
	f func(*ListOptions)
}

func (flo *funcListOption) apply(lo *ListOptions) {
	flo.f(lo)
}

func ListWithPagination(pageIndex, pageSize int) ListOption {
	return &funcListOption{f: func(options *ListOptions) {
		options.Limit = pageSize
		options.Skip = pageSize * pageIndex
	}}
}

func ListWithFilter(filteredFields ...string) ListOption {
	return &funcListOption{f: func(options *ListOptions) {
		for _, field := range filteredFields {
			options.Filter[field] = true
		}
	}}
}

func ApplyListOption(options []ListOption) *ListOptions {
	listOptions := &ListOptions{
		Filter: map[string]bool{},
	}
	for _, option := range options {
		option.apply(listOptions)
	}
	return listOptions
}

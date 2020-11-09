package pagination

type Pagination struct {
	Skip  *int64 `json:"skip,omitempty" schema:"skip,omitempty" url:"skip,omitempty"`
	Limit *int64 `json:"limit,omitempty" schema:"limit,omitempty" url:"limit,omitempty"`
}

func (page *Pagination) GetSkip() int64 {
	if page.Skip == nil {
		page.Skip = new(int64)
		*page.Skip = 0
	}
	return *page.Skip
}

func (page *Pagination) GetLimit() int64 {
	if page.Limit == nil {
		page.Limit = new(int64)
		*page.Limit = 10
	}
	return *page.Limit
}

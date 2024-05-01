package streaming

// Iterator is an interface for iterating over a collection of items.
// It handles pagination automatically behind the scenes.
type Iterator[T any] struct {
	index          int
	items          []T
	hasMorePages   bool
	nextPageCursor *string
	fetcher        func(string) ([]T, bool, *string, error)
	err            error
	maxPage        int
	currentPage    int
}

// Next returns true if there is a next item.
// It returns false when there are no more items, page limit is reached, or an error occurred.
// Call Err() to check if an error occurred.
// An iteration starts with a call to Next() and then a call to Get().
func (p *Iterator[T]) Next() bool {
	if p.index < len(p.items)-1 {
		p.index++
		return true
	}
	if !p.hasMorePages || (p.maxPage > 0 && p.currentPage >= p.maxPage) {
		return false
	}
	items, hasMorePages, nextPageCursor, err := p.fetcher(*p.nextPageCursor)
	if err != nil {
		p.err = err
		return false
	}
	p.items = items
	p.hasMorePages = hasMorePages
	p.nextPageCursor = nextPageCursor
	p.index = 0
	p.currentPage++
	if len(p.items) == 0 {
		return false
	}
	return true
}

// Get returns the current item.
// Call Next() to move to the next item.
// You should not call Get() before calling Next() at least once.
func (p *Iterator[T]) Get() T {
	return p.items[p.index]
}

// Err returns the error if an error occurred during iteration.
func (p *Iterator[T]) Err() error {
	return p.err
}

// SearchShowsByFiltersExecuteWithAutoPagination executes the request and returns a paginator to iterate over the results page by page.
// If maxPages is 0 or less, all pages will be fetched.
func (a *ShowsAPIService) SearchShowsByFiltersExecuteWithAutoPagination(r ApiSearchShowsByFiltersRequest, maxPages int) (*Iterator[Show], error) {
	res, _, err := a.SearchShowsByFiltersExecute(r)
	if err != nil {
		return nil, err
	}
	return &Iterator[Show]{
		items:          res.Shows,
		hasMorePages:   res.HasMore,
		nextPageCursor: res.NextCursor,
		fetcher: func(cursor string) ([]Show, bool, *string, error) {
			r := r.Cursor(cursor)
			res, _, err := a.SearchShowsByFiltersExecute(r)
			if err != nil {
				return nil, false, nil, err
			}
			return res.Shows, res.HasMore, res.NextCursor, nil
		},
		maxPage:     maxPages,
		currentPage: 1,
		index:       -1,
	}, nil
}

// ExecuteWithAutoPagination executes the request and returns a paginator to iterate over the results page by page.
// If maxPages is 0 or less, all pages will be fetched.
func (r ApiSearchShowsByFiltersRequest) ExecuteWithAutoPagination(maxPages int) (*Iterator[Show], error) {
	return r.ApiService.SearchShowsByFiltersExecuteWithAutoPagination(r, maxPages)
}

// ChangeWithShow represents a Change with the corresponding Show
type ChangeWithShow struct {
	Change

	// Show that the change is related to
	Show Show
}

// GetChangesExecuteWithAutoPagination executes the request and returns a paginator to iterate over the results page by page.
// Unlike Execute, type of the returned struct is ChangeWithShow instead of Change.
// If maxPages is 0 or less, all pages will be fetched.
func (a *ChangesAPIService) GetChangesExecuteWithAutoPagination(r ApiGetChangesRequest, maxPages int) (*Iterator[ChangeWithShow], error) {
	res, _, err := a.GetChangesExecute(r)
	if err != nil {
		return nil, err
	}
	return &Iterator[ChangeWithShow]{
		items:          getChangeWithShows(res),
		hasMorePages:   res.HasMore,
		nextPageCursor: res.NextCursor,
		fetcher: func(cursor string) ([]ChangeWithShow, bool, *string, error) {
			r := r.Cursor(cursor)
			res, _, err := a.GetChangesExecute(r)
			if err != nil {
				return nil, false, nil, err
			}
			return getChangeWithShows(res), res.HasMore, res.NextCursor, nil
		},
		maxPage:     maxPages,
		currentPage: 1,
		index:       -1,
	}, nil
}

// ExecuteWithAutoPagination executes the request and returns a paginator to iterate over the results page by page.
// Unlike Execute, type of the returned struct is ChangeWithShow instead of Change.
// If maxPages is 0 or less, all pages will be fetched.
func (r ApiGetChangesRequest) ExecuteWithAutoPagination(maxPages int) (*Iterator[ChangeWithShow], error) {
	return r.ApiService.GetChangesExecuteWithAutoPagination(r, maxPages)
}

func getChangeWithShows(result *ChangesResult) []ChangeWithShow {
	if result == nil {
		return nil
	}
	res := make([]ChangeWithShow, 0, len(result.Changes))
	for _, change := range result.Changes {
		res = append(res, ChangeWithShow{
			Change: change,
			Show:   result.Shows[change.ShowId],
		})
	}
	return res
}

package myfilter

import (
	"github.com/zalando/skipper/filters"
)

const (
	name         = "myfilterToken"
	defaultToken = "dalongrong"
)

type basicSpec struct {
}

type filter struct {
	allowTokens []string
}

func NewDefaultToken() filters.Spec {
	return &basicSpec{}
}

func (a filter) Response(ctx filters.FilterContext) {
	if len(a.allowTokens) == 0 {
		ctx.Response().Header.Add(defaultToken, "default-rongfengliang")
		return
	}

	token := ctx.Request().Header.Get(defaultToken)
	if token == "" {
		return
	}
	for _, o := range a.allowTokens {
		if o == token {
			ctx.Response().Header.Add(defaultToken, o)
			return
		}
	}
}

// Request is a noop
func (a filter) Request(filters.FilterContext) {}

// CreateFilter takes an optional string array.
// If any argument is not a string, it will return an error
func (spec basicSpec) CreateFilter(args []interface{}) (filters.Filter, error) {
	f := &filter{}
	for _, a := range args {
		if s, ok := a.(string); ok {
			f.allowTokens = append(f.allowTokens, s)
		} else {
			return nil, filters.ErrInvalidFilterParameters
		}
	}
	return f, nil
}

func (spec basicSpec) Name() string { return name }

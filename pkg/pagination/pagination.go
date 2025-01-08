package pagination

import (
	"reflect"
)

type (
	Request struct {
		Limit  int64 `form:"limit,default=5"`
		Offset int64 `form:"offset,default=0"`
	}

	Response struct {
		Total  int64 `json:"total"`
		Offset int64 `json:"offset"`
		Limit  int64 `json:"limit"`
	}
)

func (p Request) GetLimit() int64 {
	if p.Limit < 1 {
		return 5
	}
	return p.Limit
}

// Offset is to get offset query
func (p Request) GetOffset() int64 {
	if p.Offset < 1 {
		return 0
	}

	return p.Offset
}

func NewResponse(data interface{}, pagination Request) *Response {
	val := reflect.ValueOf(data)
	total := int64(val.Len())

	return &Response{
		Total:  total,
		Offset: pagination.GetOffset(),
		Limit:  pagination.GetLimit(),
	}
}

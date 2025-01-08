package pagination

import (
	"math"
	"reflect"
)

type (
	Request struct {
		Limit  int64 `form:"limit,default=10"`
		Offset int64 `form:"offset,default=0"`
	}

	Response struct {
		Total       int64       `json:"total"`
		CurrentPage int64       `json:"currentPage"`
		TotalPages  int64       `json:"totalPage"`
		Limit       int64       `json:"limit"`
		Data        interface{} `json:"data"`
	}
)

func NewResponse(data interface{}, pagination Request) *Response {
	// Use reflection to get the length of data
	val := reflect.ValueOf(data)
	total := int64(val.Len())

	totalPages := int64(math.Ceil(float64(total) / float64(pagination.Limit)))
	page := int64(pagination.Offset/pagination.Limit + 1)

	return &Response{
		Data:        data,
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       pagination.Limit,
	}
}

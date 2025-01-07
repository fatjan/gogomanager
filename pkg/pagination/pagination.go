package pagination

type (
	Request struct {
		Limit  int `form:"limit,default=10"`
		Offset int `form:"offset,default=0"`
	}

	Response struct {
		Total       int64       `json:"total"`
		CurrentPage int         `json:"current_page"`
		TotalPages  int         `json:"total_pages"`
		Limit       int         `json:"limit"`
		Data        interface{} `json:"data"`
	}
)

func NewResponse(data interface{}, total int64, pagination Request) Response {
	totalPages := int(total) / pagination.Limit
	if int(total)%pagination.Limit != 0 {
		totalPages++
	}

	currentPage := (pagination.Offset / pagination.Limit) + 1

	return Response{
		Data:        data,
		Total:       total,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		Limit:       pagination.Limit,
	}
}

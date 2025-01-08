package dto

type PaginationRequest struct {
	Limit  int64 `form:"limit,default=5"`
	Offset int64 `form:"offset,default=0"`
}

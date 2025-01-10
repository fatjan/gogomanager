package dto

import "strconv"

type PaginationRequest struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
}

func (p *PaginationRequest) GetLimit() int64 {
	const defaultLimit = 5

	limit, err := strconv.Atoi(p.Limit)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	return int64(limit)
}

func (p *PaginationRequest) GetOffset() int64 {
	const defaultOffset = 0

	offset, err := strconv.Atoi(p.Offset)
	if err != nil || offset < 0 {
		offset = defaultOffset
	}

	return int64(offset)
}

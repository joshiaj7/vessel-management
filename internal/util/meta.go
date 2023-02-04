package util

type OffsetPagination struct {
	Limit  int
	Offset int
	Total  int
}

func NewOffsetPagination(limit int, offset int, total int) *OffsetPagination {
	return &OffsetPagination{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

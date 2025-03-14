package utils

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoPagination struct {
	limit int
	page  int
}

func NewMongoPagination(limit int, page int) *mongoPagination {
	return &mongoPagination{
		limit: limit,
		page:  page,
	}
}
func (p *mongoPagination) GetSkip() int {
	return (p.page - 1) * p.limit
}

func (p *mongoPagination) GetLimit() int {
	return p.limit
}

func (p *mongoPagination) GetPaginationOptions() *options.FindOptions {
	return options.Find().SetSkip(int64(p.GetSkip())).SetLimit(int64(p.GetLimit()))
}

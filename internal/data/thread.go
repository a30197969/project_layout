package data

import (
	"context"
	"project_layout/internal/biz"
)

type threadRepo struct {
	data *Mysql
}

func NewThreadRepo(data *Mysql) biz.ThreadRepo {
	return &threadRepo{
		data: data,
	}
}

func (t *threadRepo) CreateThread(ctx context.Context, greeter *biz.Greeter) error {
	panic("implement me")
}

func (t *threadRepo) UpdateThread(ctx context.Context, greeter *biz.Greeter) error {
	panic("implement me")
}

func (t *threadRepo) DeleteThread(ctx context.Context, greeter *biz.Greeter) error {
	panic("implement me")
}

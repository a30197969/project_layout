package biz

import "context"

type Greeter struct {
	Hello string
}
type ThreadRepo interface {
	CreateThread(context.Context, *Greeter) error
	UpdateThread(context.Context, *Greeter) error
	DeleteThread(context.Context, *Greeter) error
}

type ThreadUsecase struct {
	repo ThreadRepo
}

func NewThreadUsecase(repo ThreadRepo) *ThreadUsecase {
	return &ThreadUsecase{
		repo: repo,
	}
}

func (t *ThreadUsecase) CreateThread(ctx context.Context, g *interface{}) error {
	panic("implement me")
}

func (t *ThreadUsecase) UpdateThread(ctx context.Context, g *interface{}) error {
	panic("implement me")
}

func (t *ThreadUsecase) DeleteThread(ctx context.Context, g *interface{}) error {
	panic("implement me")
}

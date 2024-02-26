package user

import (
	"context"

	"github.com/ronannnn/auth/cfg"
	"github.com/ronannnn/infra/models"
	"github.com/ronannnn/infra/models/request/query"
	"github.com/ronannnn/infra/models/response"
)

type Service interface {
	Create(context.Context, *models.User) error
	Update(ctx context.Context, partialUpdatedModel *models.User) (models.User, error)
	DeleteById(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	List(ctx context.Context, query query.UserQuery) (response.PageResult, error)
	GetById(ctx context.Context, id uint) (models.User, error)
}

func ProvideService(
	cfg *cfg.User,
	store Store,
) Service {
	return &ServiceImpl{
		cfg:   cfg,
		store: store,
	}
}

type ServiceImpl struct {
	cfg   *cfg.User
	store Store
}

func (srv *ServiceImpl) Create(ctx context.Context, model *models.User) (err error) {
	if model.Password == nil {
		model.Password = &srv.cfg.DefaultHashedPassword
	}
	return srv.store.Create(model)
}

func (srv *ServiceImpl) Update(ctx context.Context, partialUpdatedModel *models.User) (models.User, error) {
	return srv.store.Update(partialUpdatedModel)
}

func (srv *ServiceImpl) DeleteById(ctx context.Context, id uint) error {
	return srv.store.DeleteById(id)
}

func (srv *ServiceImpl) DeleteByIds(ctx context.Context, ids []uint) error {
	return srv.store.DeleteByIds(ids)
}

func (srv *ServiceImpl) List(ctx context.Context, query query.UserQuery) (response.PageResult, error) {
	return srv.store.List(query)
}

func (srv *ServiceImpl) GetById(ctx context.Context, id uint) (models.User, error) {
	return srv.store.GetById(id)
}

package company

import (
	"github.com/ronannnn/infra/models/response"
)

type Service interface {
	Create(model *Company) error
	Update(partialUpdatedModel *Company) (Company, error)
	DeleteById(id uint) error
	DeleteByIds(ids []uint) error
	List(query CompanyQuery) (response.PageResult, error)
	GetById(id uint) (Company, error)
}

func ProvideService(
	store Store,
) Service {
	return &ServiceImpl{
		store: store,
	}
}

type ServiceImpl struct {
	store Store
}

func (srv *ServiceImpl) Create(model *Company) (err error) {
	return srv.store.create(model)
}

func (srv *ServiceImpl) Update(partialUpdatedModel *Company) (Company, error) {
	return srv.store.update(partialUpdatedModel)
}

func (srv *ServiceImpl) DeleteById(id uint) error {
	return srv.store.deleteById(id)
}

func (srv *ServiceImpl) DeleteByIds(ids []uint) error {
	return srv.store.deleteByIds(ids)
}

func (srv *ServiceImpl) List(query CompanyQuery) (response.PageResult, error) {
	return srv.store.list(query)
}

func (srv *ServiceImpl) GetById(id uint) (Company, error) {
	return srv.store.getById(id)
}

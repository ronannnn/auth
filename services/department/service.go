package department

import (
	"github.com/ronannnn/infra/models/response"
)

type Service interface {
	Create(model *Department) error
	Update(partialUpdatedModel *Department) (Department, error)
	DeleteById(id uint) error
	DeleteByIds(ids []uint) error
	List(query DepartmentQuery) (response.PageResult, error)
	GetById(id uint) (Department, error)
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

func (srv *ServiceImpl) Create(model *Department) (err error) {
	return srv.store.create(model)
}

func (srv *ServiceImpl) Update(partialUpdatedModel *Department) (Department, error) {
	return srv.store.update(partialUpdatedModel)
}

func (srv *ServiceImpl) DeleteById(id uint) error {
	return srv.store.deleteById(id)
}

func (srv *ServiceImpl) DeleteByIds(ids []uint) error {
	return srv.store.deleteByIds(ids)
}

func (srv *ServiceImpl) List(query DepartmentQuery) (response.PageResult, error) {
	return srv.store.list(query)
}

func (srv *ServiceImpl) GetById(id uint) (Department, error) {
	return srv.store.getById(id)
}

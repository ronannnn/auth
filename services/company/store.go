package company

import (
	"github.com/ronannnn/infra/models"
	"github.com/ronannnn/infra/models/request/query"
	"github.com/ronannnn/infra/models/response"
	"gorm.io/gorm"
)

type Store interface {
	create(model *Company) error
	update(partialUpdatedModel *Company) (Company, error)
	deleteById(id uint) error
	deleteByIds(ids []uint) error
	list(query CompanyQuery) (response.PageResult, error)
	getById(id uint) (Company, error)
}

func ProvideStore(db *gorm.DB) Store {
	return StoreImpl{db: db}
}

type StoreImpl struct {
	db *gorm.DB
}

func (s StoreImpl) create(model *Company) error {
	return s.db.Create(model).Error
}

func (s StoreImpl) update(partialUpdatedModel *Company) (updatedModel Company, err error) {
	if partialUpdatedModel.Id == 0 {
		return updatedModel, models.ErrUpdatedId
	}
	if err = s.db.Updates(partialUpdatedModel).Error; err != nil {
		return
	}
	return s.getById(partialUpdatedModel.Id)
}

func (s StoreImpl) deleteById(id uint) error {
	return s.db.Delete(&Company{}, "id = ?", id).Error
}

func (s StoreImpl) deleteByIds(ids []uint) error {
	return s.db.Delete(&Company{}, "id IN ?", ids).Error
}

func (s StoreImpl) list(menuQuery CompanyQuery) (result response.PageResult, err error) {
	var total int64
	var list []Company
	if err = s.db.Model(&Company{}).Count(&total).Error; err != nil {
		return
	}
	if err = s.db.
		Scopes(query.MakeConditionFromQuery(menuQuery)).
		Preload("Apis").
		Find(&list).Error; err != nil {
		return
	}
	result = response.PageResult{
		List:     list,
		Total:    total,
		PageNum:  1,
		PageSize: int(total),
	}
	return
}

func (s StoreImpl) getById(id uint) (model Company, err error) {
	err = s.db.Preload("Apis").First(&model, "id = ?", id).Error
	return
}

package department

import (
	"github.com/ronannnn/infra/models"
	"github.com/ronannnn/infra/models/request/query"
	"github.com/ronannnn/infra/models/response"
	"gorm.io/gorm"
)

type Store interface {
	create(model *Department) error
	update(partialUpdatedModel *Department) (Department, error)
	deleteById(id uint) error
	deleteByIds(ids []uint) error
	list(query DepartmentQuery) (response.PageResult, error)
	getById(id uint) (Department, error)
}

func ProvideStore(db *gorm.DB) Store {
	return StoreImpl{db: db}
}

type StoreImpl struct {
	db *gorm.DB
}

func (s StoreImpl) create(model *Department) error {
	return s.db.Create(model).Error
}

func (s StoreImpl) update(partialUpdatedModel *Department) (updatedModel Department, err error) {
	if partialUpdatedModel.Id == 0 {
		return updatedModel, models.ErrUpdatedId
	}
	if err = s.db.Updates(partialUpdatedModel).Error; err != nil {
		return
	}
	return s.getById(partialUpdatedModel.Id)
}

func (s StoreImpl) deleteById(id uint) error {
	return s.db.Delete(&Department{}, "id = ?", id).Error
}

func (s StoreImpl) deleteByIds(ids []uint) error {
	return s.db.Delete(&Department{}, "id IN ?", ids).Error
}

func (s StoreImpl) list(menuQuery DepartmentQuery) (result response.PageResult, err error) {
	var total int64
	var list []Department
	if err = s.db.Model(&Department{}).Count(&total).Error; err != nil {
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

func (s StoreImpl) getById(id uint) (model Department, err error) {
	err = s.db.Preload("Apis").First(&model, "id = ?", id).Error
	return
}

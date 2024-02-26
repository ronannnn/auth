package department

import (
	"github.com/ronannnn/auth/services/company"
	"github.com/ronannnn/infra/models"
)

type Department struct {
	models.Base
	Name      *string          `json:"name"`
	CompanyId *uint            `json:"companyId"`
	Company   *company.Company `json:"company"`
	LeaderId  *uint            `json:"leaderId"`
	Leader    *models.User     `json:"leader"`
	ParentId  *uint            `json:"parentId"`
	Parent    *Department      `json:"parent"`
}

type DepartmentQuery struct {
	WhereQuery DepartmentWhereQuery `json:"whereQuery"`
	OrderQuery DepartmentOrderQuery `json:"orderQuery"`
}

type DepartmentWhereQuery struct {
	Name string `json:"name" search:"type:like;column:name"`
}

type DepartmentOrderQuery struct {
	CreatedAt string `json:"createdAt" search:"type:order;column:created_at"`
}

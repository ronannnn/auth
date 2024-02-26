package internal

import (
	"github.com/ronannnn/auth/services/jwt/refreshtoken"
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"github.com/ronannnn/infra/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var tables = []any{
	refreshtoken.RefreshToken{},
	models.User{},
	models.Menu{},
	models.Role{},
}

func ProvideDb(
	cfg *cfg.Db,
	log *zap.SugaredLogger,
) (db *gorm.DB, err error) {
	if db, err = infra.NewDb(cfg, false, tables); err != nil {
		return
	}
	log.Info("db initialized")
	return
}

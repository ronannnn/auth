package login

import (
	"context"

	"github.com/ronannnn/auth/services/jwt"
	"github.com/ronannnn/auth/services/jwt/refreshtoken"
	"github.com/ronannnn/auth/services/user"
	"github.com/ronannnn/infra/models"
	"gorm.io/gorm"
)

type Service interface {
	LoginByUsername(ctx context.Context, username, password string) (*Result, error)
	ChangePwd(ctx context.Context, cmd ChangeUserPwdCmd) error
}

func ProvideService(
	store user.Store,
	jwtService jwt.Service,
) Service {
	return &ServiceImpl{
		store:      store,
		jwtService: jwtService,
	}
}

type ServiceImpl struct {
	store      user.Store
	jwtService jwt.Service
}

func (srv *ServiceImpl) LoginByUsername(ctx context.Context, username, password string) (resp *Result, err error) {
	var user models.User
	if user, err = srv.store.GetByUsername(username); err == gorm.ErrRecordNotFound {
		return nil, models.ErrWrongUsernameOrPassword
	} else if err != nil {
		return
	}
	if !CheckPassword(*user.Password, password) {
		return nil, models.ErrWrongUsernameOrPassword
	}
	var refreshToken, accessToken string
	if refreshToken, accessToken, err = srv.jwtService.GenerateTokens(ctx, refreshtoken.BaseClaims{
		UserId:   user.Id,
		Username: *user.Username,
	}); err != nil {
		return
	}
	return &Result{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, err
}

func (srv *ServiceImpl) ChangePwd(ctx context.Context, cmd ChangeUserPwdCmd) (err error) {
	var user models.User
	if user, err = srv.store.GetById(cmd.UserId); err != nil {
		return
	}
	if !CheckPassword(*user.Password, cmd.OldPwd) {
		return models.ErrWrongUsernameOrPassword
	}
	var hashedNewPwd string
	if hashedNewPwd, err = HashPassword(cmd.NewPwd); err != nil {
		return
	}
	return srv.store.ChangePwd(cmd.UserId, hashedNewPwd)
}

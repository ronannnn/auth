package main

import (
	"github.com/ronannnn/auth/internal"
	"github.com/ronannnn/auth/internal/apis"
	"github.com/ronannnn/auth/services/jwt"
	"github.com/ronannnn/auth/services/jwt/accesstoken"
	"github.com/ronannnn/auth/services/jwt/refreshtoken"
	"github.com/ronannnn/auth/services/login"
	"github.com/ronannnn/auth/services/user"

	"github.com/google/wire"
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/services/apirecord"
)

var wireSet = wire.NewSet(
	// configs
	internal.ProvideSysCfg,
	internal.ProvideLogCfg,
	internal.ProvideDbCfg,
	internal.ProvideAuthCfg,
	internal.ProvideUserCfg,
	// infra
	infra.NewLog,
	infra.ProvideCasbinEnforcer,
	internal.ProvideDb,
	// middleware
	infra.ProvideMiddleware,
	// services
	apirecord.ProvideService,
	accesstoken.ProvideService,
	refreshtoken.ProvideService,
	login.ProvideService,
	jwt.ProvideService,
	user.ProvideService,
	// store
	apirecord.ProvideStore,
	refreshtoken.ProvideStore,
	user.ProvideStore,
	// server
	apis.NewHttpServer,
)

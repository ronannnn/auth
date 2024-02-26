package apis

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ronannnn/auth/services/auth"
	"github.com/ronannnn/auth/services/jwt"
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"go.uber.org/zap"
)

type HttpServer struct {
	infra.BaseHttpServer
	infraMiddleware infra.Middleware
	authService     auth.Service
	jwtService      jwt.Service
}

func NewHttpServer(
	sys *cfg.Sys,
	log *zap.SugaredLogger,
	infraMiddleware infra.Middleware,
	authService auth.Service,
	jwtService jwt.Service,
) *HttpServer {
	hs := &HttpServer{
		BaseHttpServer: infra.BaseHttpServer{
			Sys: sys,
			Log: log,
		},
		infraMiddleware: infraMiddleware,
		authService:     authService,
		jwtService:      jwtService,
	}
	// golang abstract class reference: https://adrianwit.medium.com/abstract-class-reinvented-with-go-4a7326525034
	hs.BaseHttpServer.HttpServerRunner.HttpServerBaseRunner = hs
	return hs
}

func (hs *HttpServer) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(hs.infraMiddleware.Lang)
	r.Use(hs.infraMiddleware.ReqRecorder)
	r.Use(middleware.Recoverer)
	apiV1Router := chi.NewRouter()
	apiV1Router.Mount("/auth", hs.authRouter())
	r.Mount("/api/v1", apiV1Router)
	return r
}

func (hs *HttpServer) authRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/login/username", hs.LoginByUsername)
	r.Post("/refresh", hs.RefreshTokens)
	return r
}

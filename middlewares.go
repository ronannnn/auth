package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/ronannnn/auth/services/jwt/accesstoken"
	"github.com/ronannnn/infra/cfg"
	"github.com/ronannnn/infra/models"
	"github.com/ronannnn/infra/models/response"
	"go.uber.org/zap"
)

type Middleware interface {
	// auth handlers
	AuthHandlers() []func(http.Handler) http.Handler
	Verifier(http.Handler) http.Handler
	Authenticator(http.Handler) http.Handler
	AuthInfoSetter(next http.Handler) http.Handler
}

func ProvideMiddleware(
	cfg *cfg.Auth,
	log *zap.SugaredLogger,
	accesstokenService accesstoken.Service,
) Middleware {
	return &MiddlewareImpl{
		cfg:                cfg,
		log:                log,
		accesstokenService: accesstokenService,
	}
}

type MiddlewareImpl struct {
	cfg                *cfg.Auth
	log                *zap.SugaredLogger
	accesstokenService accesstoken.Service
}

func (m *MiddlewareImpl) AuthHandlers() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		m.Verifier,
		m.Authenticator,
		m.AuthInfoSetter,
	}
}

func (m *MiddlewareImpl) Verifier(next http.Handler) http.Handler {
	if !m.cfg.Enabled {
		return next
	}
	return jwtauth.Verifier(m.accesstokenService.GetJwtAuth())(next)
}

// Authenticator override chi.Authenticator
func (m *MiddlewareImpl) Authenticator(next http.Handler) http.Handler {
	if !m.cfg.Enabled {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			response.ErrAccessToken(w, r, err)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			response.ErrAccessToken(w, r, fmt.Errorf("invalid token"))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

// AuthInfoSetter is a middleware that sets the auth info(user id and username) for the request.
// It must be placed after jwt middleware.
func (m *MiddlewareImpl) AuthInfoSetter(next http.Handler) http.Handler {
	if !m.cfg.Enabled {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())
		username, _ := token.Get("username")
		userId, _ := token.Get("userId")
		ctx := context.WithValue(r.Context(), models.CtxKeyUserId, uint(userId.(float64)))
		ctx = context.WithValue(ctx, models.CtxKeyUsername, username.(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

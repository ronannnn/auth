package apis

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ronannnn/auth/services/auth"
	"github.com/ronannnn/infra/models/response"
)

func (hs *HttpServer) LoginByUsername(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload auth.LoginUsernameCommand
	if err = render.DefaultDecoder(r, &payload); err != nil {
		response.FailWithErr(w, r, err)
		return
	}

	var result *auth.AuthResult
	if result, err = hs.authService.LoginByUsername(r.Context(), payload.Username, payload.Password); err != nil {
		response.FailWithErr(w, r, err)
		return
	}
	response.OkWithData(w, r, result)
}

func (hs *HttpServer) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload auth.RefreshTokensCommand
	if err = render.DefaultDecoder(r, &payload); err != nil {
		response.ErrRefreshToken(w, r, err)
		return
	}

	var refreshToken, accessToken string
	if refreshToken, accessToken, err = hs.jwtService.UpdateTokens(r.Context(), payload.RefreshToken); err != nil {
		response.ErrRefreshToken(w, r, err)
		return
	}

	response.OkWithData(w, r, &auth.AuthResult{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}

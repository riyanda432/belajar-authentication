package route_mobile_app_v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	handler "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/handler/mobile_app"
)

func UserRouter(h handler.IUserHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/sign-up", h.Create)
	r.Post("/sign-in", h.Login)

	return r
}

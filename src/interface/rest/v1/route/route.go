package route_v1

import (
	"net/http"

	"github.com/riyanda432/belajar-authentication/src/interface/rest/middleware"
	usecases "github.com/riyanda432/belajar-authentication/src/app"
	rest_interface "github.com/riyanda432/belajar-authentication/src/interface/rest/interface"
	"github.com/riyanda432/belajar-authentication/src/interface/rest/response"
	mobile_app_handler_v1 "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/handler/mobile_app"
	mobile_app_route_v1 "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/route/mobile_app"

	"github.com/go-chi/chi/v5"
	// "github.com/riyanda432/belajar-authentication/src/interface/rest/middleware"
)

// MobileAppRouter a completely separate router for mobile app routes
func MobileAppRouterV1(
	oh mobile_app_handler_v1.IUserHandler,
) http.Handler {
	r := chi.NewRouter()

	// mobile app center header
	r.Use(middleware.CheckMobileAppHeader)

	// working day routes
	r.Mount("/users", mobile_app_route_v1.UserRouter(oh))

	// register more mobile-app routes over here ...

	return r
}

func MakeRouteV1(
	r *chi.Mux,
	logger rest_interface.ILogger,
	useCases usecases.AllUseCases,
) *chi.Mux {
	// instantiate the handlers here ...
	respClient := response.NewResponseClient()
	sa_v1_user := mobile_app_handler_v1.NewUserHandler(respClient, useCases.MobileAppV1UseCases.UserUseCase)
	r.Route("/api/v1", func(r chi.Router) {

		// mobile app routes
		r.Mount("/mobile-app", MobileAppRouterV1(sa_v1_user))

		// register more routes over here ...
	})

	return r
}

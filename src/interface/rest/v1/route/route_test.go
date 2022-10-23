package route_v1

import (
	"testing"

	"github.com/go-chi/chi/v5"
	mock_handler "github.com/riyanda432/belajar-authentication/mocks/interface/rest/v1/handler/mobile_app"
	"github.com/stretchr/testify/assert"
)

func TestMobileAppRouterV1(t *testing.T) {
	r := MobileAppRouterV1(&mock_handler.MockIUserHandler{})
	assert.IsType(t, &chi.Mux{}, r)
}


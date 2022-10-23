package usecases

import (
	"testing"

	mock_repo_mobile_app "github.com/riyanda432/belajar-authentication/mocks/domain/repositories/mobile_app"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := NewAllUsecase(
		&mock_repo_mobile_app.MockUserRepository{},
	)
	assert.IsType(t, AllUseCases{}, d)
}

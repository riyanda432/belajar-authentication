package postgres

import (
	"context"
	"errors"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	repositories "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"
	infra_model "github.com/riyanda432/belajar-authentication/src/infra/models"

	"gorm.io/gorm"
)

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		connection: db,
	}
}

var _ repositories.UserRepository = &userRepository{}

func (repo *userRepository) Persist(
	ctx context.Context,
	user *entities.User,
) error {
	if user == nil {
		return errors.New("user entity can not empty")
	}
	tx := repo.connection.WithContext(ctx).Begin()
	model := infra_model.ToUserModel(*user)
	resCreate := tx.Session(&gorm.Session{
		PrepareStmt:          true,
		QueryFields:          true,
		FullSaveAssociations: true,
	}).Save(&model)
	if resCreate.Error != nil {
		tx.Rollback()
		return resCreate.Error
	}

	tx.Commit()
	ose := model.ToEntity()

	*user = *ose
	return nil
}

func (repo *userRepository) DetailByUserName(
	ctx context.Context,
	UserName string,
) (*entities.User, error) {
	var model *infra_model.User
	var entity *entities.User
	filter := infra_model.User{
		UserName: UserName,
	}

	if err := repo.connection.WithContext(ctx).
		Session(&gorm.Session{PrepareStmt: true, QueryFields: true}).
		Where(filter).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	entity = model.ToEntity()
	return entity, nil
}

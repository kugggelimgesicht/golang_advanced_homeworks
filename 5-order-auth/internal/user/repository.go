package user

import (
	"validation-api/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository { return &UserRepository{Database: database} }

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo *UserRepository) FindBySessionId(sessionId string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

package userstorage

import (
	"context"
	"demo/common"
	"demo/modules/user/usermodel"
)

func (s *sqlStore) UpdateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

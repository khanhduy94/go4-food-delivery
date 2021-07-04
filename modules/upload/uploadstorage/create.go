package uploadstorage

import (
	"context"
	"demo/common"
	"demo/modules/upload/uploadmodel"
)

func (store *sqlStore) CreateImage(context context.Context, data *common.Image) error {
	db := store.db

	if err := db.Table(uploadmodel.Upload{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

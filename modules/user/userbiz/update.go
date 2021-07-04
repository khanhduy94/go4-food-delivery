package userbiz

import (
	"context"
	"demo/modules/user/usermodel"
)

type UserStorage interface {
	UpdateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

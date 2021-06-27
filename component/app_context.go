package component

import (
	"demo/component/upload"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() upload.Provider
	GetSecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider upload.Provider
	secretKey  string
}

func NewAppContext(db *gorm.DB, upProvider upload.Provider, secretKey string) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetUploadProvider() upload.Provider {
	return ctx.upProvider
}

func (ctx *appCtx) GetSecretKey() string { return ctx.secretKey }

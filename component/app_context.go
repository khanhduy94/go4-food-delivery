package component

import (
	"demo/component/upload"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() *upload.Provider
}

type appCtx struct {
	db         *gorm.DB
	upProvider upload.Provider
}

func NewAppContext(db *gorm.DB, upProvider upload.Provider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetUploadProvider() upload.Provider {
	return ctx.upProvider
}

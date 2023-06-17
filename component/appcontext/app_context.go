package appcontext

import (
	"food-delivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) *appContext {
	return &appContext{
		db:             db,
		uploadProvider: provider,
		secretKey:      secretKey,
	}
}
func (ctx *appContext) GetMainDBConnection() *gorm.DB                 { return ctx.db }
func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }
func (ctx *appContext) SecretKey() string                             { return ctx.secretKey }

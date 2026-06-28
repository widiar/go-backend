package handlers

import (
	"backendmaw/services"

	"go.mau.fi/whatsmeow/store/sqlstore"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth     *AuthHandler
	Banner   *BannerHandler
	Feature  *FeatureHandler
	Merchant *MerchantHandler
	Wa       *WaHandler
}

func Setup(db *gorm.DB, waContainer *sqlstore.Container) *Handlers {
	authService := services.NewAuthService(db)
	bannerService := services.NewBannerService(db)
	featureService := services.NewFeatureService(db)
	merchantService := services.NewMerchantService(db)
	waService := services.NewWaService(waContainer)

	return &Handlers{
		Auth:     NewAuthHandler(authService),
		Banner:   NewBannerHandler(bannerService),
		Feature:  NewFeatureHandler(featureService),
		Merchant: NewMerchantHandler(merchantService),
		Wa:       NewWaHandler(waService),
	}
}

package dto

type MerchantResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MerchantRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type MerchantFeatureResponse struct {
	MerchantResponse
	Features []FeatureResponse `json:"features"`
}

type MerchantFeatureRequest struct {
	Items []MerchantFeatureItemReqeust `json:"items" validate:"required,min=1,dive"`
}

type MerchantFeatureItemReqeust struct {
	MerchantId string   `json:"merchant_id" validate:"required,uuid"`
	FeatureId  []string `json:"feature_id" validate:"required,dive,uuid"`
}

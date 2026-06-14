package dto

type FeatureRequest struct {
	Name  string `json:"name" validate:"required"`
	Label string `json:"label" validate:"required"`
}

type FeatureResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

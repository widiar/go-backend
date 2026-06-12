package dto

type FeatureRequest struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type FeatureResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

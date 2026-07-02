package dto

type WfhRequest struct {
	Start string `query:"start" validate:"required"`
	End   string `query:"end" validate:"required"`
}

type WfhResponse struct {
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type CalendarConfigRequest struct {
	Type  string `json:"type" validate:"required"`
	Value string `json:"value" validate:"required"`
}

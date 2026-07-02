package dto

type WfhRequest struct {
	Start string `query:"start" validate:"required"`
	End   string `query:"end" validate:"required"`
}

package dto

type SendWaMessageRequest struct {
	Sender  string `json:"sender" validate:"required"`
	Target  string `json:"target" validate:"required"`
	Message string `json:"message" validate:"required"`
}

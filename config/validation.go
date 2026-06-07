package config

import (
	"backendmaw/dto"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func SetupHttpErrorHandler(c *echo.Context, err error) {
	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}
	if valErr, ok := err.(validator.ValidationErrors); ok {
		c.Logger().Warn("ERROR !")
		var validationErr []map[string]string
		for _, e := range valErr {
			validationErr = append(validationErr, map[string]string{
				"field":   strcase.ToSnake(e.Field()),
				"message": fmt.Sprintf("Failed on '%s' validation", e.Tag()),
			})
		}
		_ = c.JSON(http.StatusBadRequest, dto.ResponseDto{
			Status:      http.StatusBadRequest,
			ErrorSchema: "Validation Error",
			Payload:     nil,
			Validation:  validationErr,
		})
		return
	}
	if he, ok := err.(*echo.HTTPError); ok {
		_ = c.JSON(he.Code, dto.ResponseDto{
			Status:      he.Code,
			ErrorSchema: fmt.Sprintf("%v", he.Message),
			Payload:     nil,
			Validation:  nil,
		})
		return
	}
	_ = c.JSON(http.StatusInternalServerError, dto.ResponseDto{
		Status:      http.StatusInternalServerError,
		ErrorSchema: "Internal Server Error",
		Payload:     nil,
		Validation:  nil,
	})
}

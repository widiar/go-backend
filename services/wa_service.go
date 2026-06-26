package services

import (
	"backendmaw/config"
	"backendmaw/dto"
	"context"
	"net/http"
	"time"
)

func LoginWa() (*dto.ResponseDto, error) {
	if config.WA.IsConnected() {
		return new(dto.SuccessResponse("Already connected")), nil
	}

	qrChan, _ := config.WA.GetQRChannel(context.Background())
	if err := config.WA.Connect(); err != nil {
		return new(dto.ErrorResponse()), err
	}
	timeout := time.After(20 * time.Second)
	for {
		select {
		case evt := <-qrChan:
			if evt.Event == "code" {
				return new(dto.SuccessResponse(evt.Code)), nil
			}
			return new(dto.SuccessResponse(evt.Event)), nil
		case <-timeout:
			return new(dto.FailedResponse("Timeout", http.StatusRequestTimeout)), nil
		}
	}
}

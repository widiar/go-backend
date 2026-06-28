package services

import (
	"backendmaw/dto"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WaService struct {
	WA *sqlstore.Container
}

func NewWaService(wa *sqlstore.Container) *WaService {
	return &WaService{WA: wa}
}

func (s *WaService) Login() (*dto.ResponseDto, error) {
	device := s.WA.NewDevice()
	clientLog := waLog.Stdout("Client", "WARN", true)
	client := whatsmeow.NewClient(device, clientLog)

	qrChan, _ := client.GetQRChannel(context.Background())
	err := client.Connect()
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	go func() {
		for evt := range qrChan {
			if evt.Event == "success" {
				fmt.Printf("\n[BERHASIL] Nomor %s berhasil login!\n", client.Store.ID.User)
				time.Sleep(60 * time.Second)
				client.Disconnect()
				return
			} else if evt.Event == "timeout" {
				fmt.Println("\n[TIMEOUT] QR Code kedaluwarsa.")
				client.Disconnect()
				return
			}
		}
	}()

	timeout := time.After(40 * time.Second)
	select {
	case evt := <-qrChan:
		if evt.Event == "code" {
			return new(dto.SuccessResponse(evt.Code)), nil
		}
	case <-timeout:
		return new(dto.FailedResponse("Failed get QR", http.StatusRequestTimeout)), nil
	}
	return new(dto.ErrorResponse()), nil
}

func (s *WaService) SendMessage(req *dto.SendWaMessageRequest) (*dto.ResponseDto, error) {
	devices, err := s.WA.GetAllDevices(context.Background())
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	var device *store.Device
	for _, d := range devices {
		if d.ID.User == req.Sender {
			device = d
			break
		}
	}
	if device == nil {
		return new(dto.ErrorResponse()), fmt.Errorf("no device with user %s", req.Sender)
	}

	clientLog := waLog.Stdout("Client", "WARN", true)
	client := whatsmeow.NewClient(device, clientLog)
	err = client.Connect()
	if err != nil {
		client.Disconnect()
		return new(dto.ErrorResponse()), err
	}
	go func() {
		time.Sleep(20 * time.Second)
		client.Disconnect()
		fmt.Println("[END] SendMessage WA", "error", err)
	}()

	targetJid := types.NewJID(req.Target, types.DefaultUserServer)
	msg := &waE2E.Message{
		Conversation: new(req.Message),
	}
	resp, err := client.SendMessage(context.Background(), targetJid, msg)
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(resp.ID)), nil
}

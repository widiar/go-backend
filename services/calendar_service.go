package services

import (
	"backendmaw/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"time"
)

type CalendarService struct {
}

func NewCalendarService() *CalendarService {
	return &CalendarService{}
}

func (s *CalendarService) EventWfh(dtoReq *dto.WfhRequest) (*dto.ResponseDto, error) {
	stringEscape := url2.QueryEscape("id.indonesian#holiday@group.v.calendar.google.com")
	start := url2.QueryEscape(fmt.Sprintf("%s10:00:00Z", dtoReq.Start))
	end := url2.QueryEscape(fmt.Sprintf("%s10:00:00Z", dtoReq.End))
	key := os.Getenv("GOOGLE_API_KEY")
	url := fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/%s/events?timeMax=%s&timeMin=%s&key=%s",
		stringEscape,
		end,
		start,
		key)

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return new(dto.ErrorResponse()), err
	}

	return new(dto.SuccessResponse(data)), nil
}

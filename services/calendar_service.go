package services

import (
	"backendmaw/dto"
	"backendmaw/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CalendarService struct {
	DB *gorm.DB
}

func NewCalendarService(db *gorm.DB) *CalendarService {
	return &CalendarService{DB: db}
}

func (s *CalendarService) EventWfh() (*dto.ResponseDto, error) {
	var calendars []models.ConfigCalendar
	if err := s.DB.Where("type IN ?", []string{"START_CALENDAR", "END_CALENDAR"}).Find(&calendars).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	var calendarStart, calendarEnd models.ConfigCalendar
	for _, c := range calendars {
		if c.Type == "START_CALENDAR" {
			calendarStart = c
		} else if c.Type == "END_CALENDAR" {
			calendarEnd = c
		}
	}

	stringEscape := url2.QueryEscape("id.indonesian#holiday@group.v.calendar.google.com")
	start := url2.QueryEscape(fmt.Sprintf("%sT10:00:00Z", calendarStart.Value))
	end := url2.QueryEscape(fmt.Sprintf("%sT10:00:00Z", calendarEnd.Value))
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
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := io.ReadAll(resp.Body)
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return new(dto.ErrorResponse()), err
	}
	items, ok := data["items"].([]any)
	if !ok {
		return new(dto.ErrorResponse()), fmt.Errorf("items not found")
	}

	var excludeDate []models.ConfigCalendar
	if errDB := s.DB.Where("type = ?", "EXCLUDE").Find(&excludeDate).Error; errDB != nil {
		return new(dto.ErrorResponse()), err
	}
	makeExcludeDate := make(map[string]bool)
	for _, c := range excludeDate {
		makeExcludeDate[c.Value] = true
	}

	var response []dto.WfhResponse
	liburMap := make(map[string]bool)
	for _, item := range items {
		event := item.(map[string]any)
		dateStart := event["start"].(map[string]any)["date"].(string)
		dateEnd := event["end"].(map[string]any)["date"].(string)
		if makeExcludeDate[dateStart] || makeExcludeDate[event["summary"].(string)] {
			continue
		}
		liburMap[dateStart] = true
		response = append(response, dto.WfhResponse{
			Start:       dateStart,
			End:         dateEnd,
			Description: event["summary"].(string),
			Tag:         "libur",
		})
	}

	//calculate wfh
	var startWfh []models.ConfigCalendar
	if errDB := s.DB.Where("type LIKE ?", "START_WFH%").Find(&startWfh).Error; errDB != nil {
		return new(dto.ErrorResponse()), err
	}
	if len(startWfh) < 4 {
		return new(dto.FailedResponse("wfh less than 4", http.StatusBadRequest)), fmt.Errorf("start wfh not found")
	}
	var startWfh1Str, startWfh2Str, startWfh3Str, startWfh4Str string
	for _, c := range startWfh {
		if c.Type == "START_WFH_1" {
			startWfh1Str = c.Value
		} else if c.Type == "START_WFH_2" {
			startWfh2Str = c.Value
		} else if c.Type == "START_WFH_3" {
			startWfh3Str = c.Value
		} else if c.Type == "START_WFH_4" {
			startWfh4Str = c.Value
		}
	}

	startWfh1, _ := initWfh(startWfh1Str, &response, "1")
	startWfh2, _ := initWfh(startWfh2Str, &response, "2")
	startWfh3, _ := initWfh(startWfh3Str, &response, "3")
	startWfh4, _ := initWfh(startWfh4Str, &response, "4")
	counter1 := 0
	counter2 := 0
	counter3 := 0
	counter4 := 0
	startDay, _ := time.Parse(dateFormat, calendarStart.Value)
	endDay, _ := time.Parse(dateFormat, calendarEnd.Value)

	for d := startDay; d.Before(endDay); d = d.AddDate(0, 0, 1) {
		counter1 = calculateWfh(d, startWfh1, counter1, &response, liburMap, "1")
		counter2 = calculateWfh(d, startWfh2, counter2, &response, liburMap, "2")
		counter3 = calculateWfh(d, startWfh3, counter3, &response, liburMap, "3")
		counter4 = calculateWfh(d, startWfh4, counter4, &response, liburMap, "4")
	}

	return new(dto.SuccessResponse(response)), nil
}

const dateFormat = "2006-01-02"

func initWfh(date string, evt *[]dto.WfhResponse, label string) (time.Time, error) {
	*evt = append(*evt, dto.WfhResponse{
		Start:       date,
		End:         date,
		Description: label,
		Tag:         "wfh",
	})
	return time.Parse(dateFormat, date)
}

func isWeekend(day time.Time) bool {
	return day.Weekday() == time.Saturday || day.Weekday() == time.Sunday
}
func calculateWfh(day time.Time, start time.Time, counter int, result *[]dto.WfhResponse, liburMap map[string]bool, label string) int {
	dayStr := day.Format(dateFormat)
	isValidWorkDay := !isWeekend(day) && !liburMap[dayStr]
	if day.After(start) && isValidWorkDay {
		counter++
		if counter == 4 {
			*result = append(*result, dto.WfhResponse{
				Start:       dayStr,
				End:         dayStr,
				Description: label,
				Tag:         "wfh",
			})
			counter = 0
		}
	}
	return counter
}

func (s *CalendarService) ConfigCalendar(req *dto.CalendarConfigRequest) (*dto.ResponseDto, error) {
	var model models.ConfigCalendar
	if strings.Contains(strings.ToUpper(req.Type), "START_WFH") || strings.Contains(strings.ToUpper(req.Type), "CALENDAR") {
		if err := s.DB.First(&model, "type = ?", req.Type).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.ErrorResponse()), err
		}
		model.Type = req.Type
		model.Value = req.Value
		s.DB.Save(&model)
	} else {
		model.Type = req.Type
		model.Value = req.Value
		s.DB.Create(&model)
	}

	return new(dto.SuccessResponse("OK")), nil
}

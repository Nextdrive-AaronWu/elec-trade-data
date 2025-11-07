package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Nextdrive-AaronWu/elec-trade-data/internal/model"
)

// FetchDailyData 只負責 API 呼叫，日期由 main.go 傳入
func FetchDailyData(startDate string) (*model.APIResponse, error) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("API_BASE_URL not set")
	}

	url := fmt.Sprintf("%s?startDate=%s", baseURL, startDate)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var result model.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	return &result, nil
}

package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type dataGempaResponse struct {
	Infogempa struct {
		Gempa struct {
			Tanggal     string `json:"Tanggal"`
			Jam         string `json:"Jam"`
			DateTime    string `json:"DateTime"`
			Coordinates string `json:"Coordinates"`
			Lintang     string `json:"Lintang"`
			Bujur       string `json:"Bujur"`
			Magnitude   string `json:"Magnitude"`
			Kedalaman   string `json:"Kedalaman"`
			Wilayah     string `json:"Wilayah"`
			Potensi     string `json:"Potensi"`
			Dirasakan   string `json:"Dirasakan"`
			Shakemap    string `json:"Shakemap"`
		} `json:"gempa"`
	} `json:"Infogempa"`
}

func getGempa() (*dataGempaResponse, error) {
	resp, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json")
	if err != nil {
		return nil, fmt.Errorf("failed to get earthquake data: %v", err)
	}
	defer resp.Body.Close()

	var autoGempaResp dataGempaResponse
	if err := json.NewDecoder(resp.Body).Decode(&autoGempaResp); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}
	return &autoGempaResp, nil
}

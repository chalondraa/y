package utils

import (
	"bingai-bot/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DalleRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type DalleResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func GenerateImage(prompt string) (string, error) {
	url := "https://api.openai.com/v1/images/generations"

	reqBody := DalleRequest{
		Prompt: prompt,
		N:      1,
		Size:   "512x512",
	}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+config.DalleAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var dalleRes DalleResponse
	err = json.NewDecoder(resp.Body).Decode(&dalleRes)
	if err != nil || len(dalleRes.Data) == 0 {
		return "", fmt.Errorf("gagal decode response")
	}
	return dalleRes.Data[0].URL, nil
}

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"thedekk/Shiza/internal/env"
	"bytes"

)
type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type Reasoning struct {
    Enabled bool `json:"enabled"`
}
type RequestBody struct {
    Model     string    `json:"model"`
    Messages  []Message `json:"messages"`
    Reasoning Reasoning `json:"reasoning"`
}

type Choice struct {
    Message struct {
        Content string `json:"content"`
    } `json:"message"`
}

type Response struct {
    Choices []Choice `json:"choices"`
}

func Request(msg string) (*string, error) {
	client := &http.Client{}

	jsonData, err := os.Open("internal/api/test.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil, err
	}

	defer jsonData.Close()

	var bodyJSON RequestBody
	if err := json.NewDecoder(jsonData).Decode(&bodyJSON); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	bodyJSON.Messages[0].Content += msg

	result, err := json.Marshal(bodyJSON)
	if err != nil {
	    return nil, err
	}

	req, err := http.NewRequest(
		"POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(result),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	env := env.Config{}
	env.Load()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+env.KeyAPI)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
  	  return nil, err
	}

	text := apiResp.Choices[0].Message.Content
	return &text, nil
}

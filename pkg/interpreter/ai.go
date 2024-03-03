package interpreter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Data structure for sending request to OpenAI
type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// CallGPT takes row data and calls the OpenAI API using curl, returning the interpreted response
func CallGPT(rowData string) (string, error) {
	systemMessage := "You are a smart assistant specialized in organizing and formatting event details into iCalendar (.ics) format for Google Calendar entries. Given a description of events including dates, locations, and times, your task is to extract these details and format them into an iCalendar event format. Each event entry should include properties like SUMMARY, DTSTART, DTEND, DESCRIPTION, and the TZID parameter for timezone. If specific times are not mentioned, make an educated guess based on the event context and convert all times to a 24-hour format. Use the location to infer the appropriate timezone and include this using the TZID parameter in DTSTART and DTEND. If the input does not provide enough data to format into an iCalendar event or if the data is unsuitable, respond with NULL. Note: This is an automated script sending you requests. Your responses will be directly used to generate .ics files, so please ensure accuracy and clarity in the iCalendar format. Here is an example: 'INPUT: [('3/13/2024', 'Kinkaku-ji, Kyoto', '09:00-10:00', 'Visit to Kinkaku-ji'), ('3/13/2024', 'Fushimi Inari, Kyoto', '11:00-13:00', 'Visit to Fushimi Inari'), ('3/13/2024', 'Kyoto', '13:00-14:00', 'Lunch'), ('3/13/2024', 'Gion District, Kyoto', '14:30-', 'Explore Gion District, no booking needed. Good for shopping and Maiko sightings.')]' and you output would be: 'BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nBEGIN:VTIMEZONE\nTZID:Asia/Tokyo\nBEGIN:STANDARD\nDTSTART:20240313T000000\nTZOFFSETFROM:+0900\nTZOFFSETTO:+0900\nTZNAME:JST\nEND:STANDARD\nEND:VTIMEZONE\nBEGIN:VEVENT\nDTSTART;TZID=Asia/Tokyo:20240313T090000\nDTEND;TZID=Asia/Tokyo:20240313T100000\nSUMMARY:Visit to Kinkaku-ji\nDESCRIPTION:Visit to Kinkaku-ji\nLOCATION:Kinkaku-ji, Kyoto\nEND:VEVENT\nBEGIN:VEVENT\nDTSTART;TZID=Asia/Tokyo:20240313T110000\nDTEND;TZID=Asia/Tokyo:20240313T130000\nSUMMARY:Visit to Fushimi Inari\nDESCRIPTION:Visit to Fushimi Inari\nLOCATION:Fushimi Inari, Kyoto\nEND:VEVENT\nBEGIN:VEVENT\nDTSTART;TZID=Asia/Tokyo:20240313T130000\nDTEND;TZID=Asia/Tokyo:20240313T140000\nSUMMARY:Lunch\nDESCRIPTION:Lunch\nLOCATION:Kyoto\nEND:VEVENT\nBEGIN:VEVENT\nDTSTART;TZID=Asia/Tokyo:20240313T143000\nDTEND;TZID=Asia/Tokyo:20240313T163000\nSUMMARY:Explore Gion District\nDESCRIPTION:Explore Gion District, no booking needed. Good for shopping and Maiko sightings.\nLOCATION:Gion District, Kyoto\nEND:VEVENT\nEND:VCALENDAR'"
	userMessage := fmt.Sprintf("Format this for Google Calendar: %s", rowData)

	// Prepare the payload
	payload := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "system", Content: systemMessage},
			{Role: "user", Content: userMessage},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Use os.Getenv to read the API key from an environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	cmd := exec.Command("curl", "https://api.openai.com/v1/chat/completions",
		"-H", "Content-Type: application/json",
		"-H", fmt.Sprintf("Authorization: Bearer %s", apiKey),
		"-d", string(payloadBytes),
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	var response OpenAIResponse
	if err := json.Unmarshal(out.Bytes(), &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(response.Choices) > 0 && response.Choices[0].Message.Content != "NULL" && response.Choices[0].Message.Content != "" {
		return response.Choices[0].Message.Content, nil
	}

	return "", nil
}

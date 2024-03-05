// in main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Luew2/CapitolReef/pkg/ics"
	"github.com/Luew2/CapitolReef/pkg/interpreter"
	"github.com/Luew2/CapitolReef/pkg/parser"
	"github.com/Luew2/CapitolReef/pkg/spreadsheet"
)

func main() {
	var (
		filePath string
		timezone string
	)
	flag.StringVar(&filePath, "file", "", "Path to the spreadsheet file")
	flag.StringVar(&timezone, "timezone", "UTC", "Timezone for the events (e.g., 'America/New_York')")
	flag.Parse() // Parse the flags

	// Check if the filePath is provided
	if filePath == "" {
		log.Fatal("No file path provided. Use the --file flag to specify the spreadsheet file path.")
	}

	_, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("Invalid timezone provided: %v", err)
	}

	results, err := spreadsheet.ParseSpreadsheet(filePath)
	if err != nil {
		log.Fatalf("Failed to parse spreadsheet: %v", err)
	}

	var allEvents []parser.EventDetail

	for i, rowMap := range results {
		rowDataStr := formatRowData(rowMap)
		interpretedData, err := interpreter.CallGPT(rowDataStr)
		if err != nil {
			log.Printf("Failed to interpret data for row %d: %v", i+1, err)
			continue
		}

		eventDetails, err := parser.ParseICalEvents(interpretedData)
		if err != nil {
			log.Printf("Failed to parse iCal data for row %d: %v", i+1, err)
			continue
		}

		allEvents = append(allEvents, eventDetails...)
	}

	// Turn events into ical events
	errs := ics.CreateICS(allEvents, "calendar.ics", timezone)
	if errs != nil {
		log.Fatalf("Failed to create ICS file: %v", err)
	}

}

func formatRowData(rowMap map[string]string) string {
	var rowData []string
	for key, value := range rowMap {
		rowData = append(rowData, fmt.Sprintf("%s: %s", key, value))
	}
	return strings.Join(rowData, "\n")
}

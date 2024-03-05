package parser

import (
	"strings"
)

func ParseICalEvents(icalData string) ([]EventDetail, error) {
	var events []EventDetail
	const eventStart = "BEGIN:VEVENT"
	const eventEnd = "END:VEVENT"
	eventBlocks := strings.Split(icalData, eventStart)

	for _, block := range eventBlocks[1:] { // Skip the first split, as it's before the first BEGIN:VEVENT
		eventDetail := EventDetail{}
		lines := strings.Split(block[:strings.Index(block, eventEnd)], "\n")
		for _, line := range lines {
			switch {
			case strings.HasPrefix(line, "DTSTART"):
				eventDetail.Start = strings.TrimPrefix(line, "DTSTART:")
			case strings.HasPrefix(line, "DTEND"):
				eventDetail.End = strings.TrimPrefix(line, "DTEND:")
			case strings.HasPrefix(line, "SUMMARY"):
				eventDetail.Summary = strings.TrimPrefix(line, "SUMMARY:")
			case strings.HasPrefix(line, "DESCRIPTION"):
				eventDetail.Description = strings.TrimPrefix(line, "DESCRIPTION:")
			case strings.HasPrefix(line, "LOCATION"):
				eventDetail.Location = strings.TrimPrefix(line, "LOCATION:")
			}
		}
		events = append(events, eventDetail)
	}

	return events, nil
}

package ics

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Luew2/CapitolReef/pkg/parser"
	ics "github.com/arran4/golang-ical"
)

func CreateICS(events []parser.EventDetail, filename string, timezone string) error {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)

	for _, event := range events {
		icalEvent := cal.AddEvent(event.Summary)
		icalEvent.SetStartAt(parseTime(event.Start))
		icalEvent.SetEndAt(parseTime(event.End))
		icalEvent.SetSummary(event.Summary)
		icalEvent.SetDescription(event.Description)
		icalEvent.SetLocation(event.Location)
	}

	icsString := cal.Serialize()

	icsStringProcessed := postProcessICalendarString(icsString, timezone)

	return os.WriteFile(filename, []byte(icsStringProcessed), 0644)
}

func postProcessICalendarString(icsString string, timezone string) string {
	tzid := fmt.Sprintf(";TZID=%s:", timezone)

	lines := strings.Split(icsString, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "DTSTART:") || strings.HasPrefix(line, "DTEND:") {
			runes := []rune(line)
			if runes[len(runes)-2] == 'Z' {
				runes = runes[:len(runes)-2]
			} else if runes[len(runes)-1] == 'Z' {
				runes = runes[:len(runes)-1]
			}
			// Convert runes back to string and append the timezone
			line = string(runes)

			// Replace the colon with the tzid, ensuring timezone is properly appended
			line = strings.Replace(line, ":", tzid, 1)
			lines[i] = line
		}
	}
	return strings.Join(lines, "\n")
}

func parseTime(timeStr string) time.Time {
	const layout = "20060102T150405Z"

	parsedTime, _ := time.Parse(layout, timeStr)
	return parsedTime
}

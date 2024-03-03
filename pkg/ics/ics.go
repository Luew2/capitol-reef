package ics

import (
	"os"
	"time"

	"github.com/Luew2/CapitolReef/pkg/parser"
	ics "github.com/arran4/golang-ical"
)

func CreateICS(events []parser.EventDetail, filename string) error {
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

	return os.WriteFile(filename, []byte(icsString), 0644)
}

func parseTime(timeStr string) time.Time {
	// Parse timeStr into time.Time object here
	parsedTime, _ := time.Parse("20060102T150405", timeStr)
	return parsedTime
}

// take in event details
// put them all in one big ics file
// write out ics file to . directory.
// return complete string or a true bool or something

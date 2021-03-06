package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func fetchUpcomingEvents(service *calendar.Service, count int64) {
	t := time.Now().Format(time.RFC3339)
	events, err := service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(count).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}

}

func main() {
	service := createCalService()
	fetchUpcomingEvents(service, 20)
}

package tasks_test

import (
	"Prioritized/v0/tasks"
	"fmt"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	var tests = []struct{
		str		string
		want		[7]int
		time_zone 	*time.Location
	}{
		{
			"12:00:00+00:00",
			[7]int{0, 1, 1, 12, 0, 0, 0},
			time.UTC,
		},
		{
			"10:30:30+06:00",
			[7]int{0, 1, 1, 10, 30, 30, 0},
			time.FixedZone("UTC+6", 6*60*60),
		},
	}


	for _, tt := range(tests) {
		failed := false
		timeString := tasks.Time(tt.str)
		
		parsedTime, err := timeString.Parse()

		if err != nil {
			t.Errorf("failed to parse time string: %v", err.Error())
		}

		if got := parsedTime.Year(); got != tt.want[0] {
			t.Errorf("year: want %v got %v for input %v", tt.want[0], got, tt.str)
			failed = true
		}
		if got := parsedTime.Month(); int(got) != tt.want[1] {
			t.Errorf("month: want %v got %v for input %v", tt.want[1], got, tt.str)
			failed = true
		}
		if got := parsedTime.Day(); got != tt.want[2] {
			t.Errorf("day: want %v got %v for input %v", tt.want[2], got, tt.str)
			failed = true
		}
		if got := parsedTime.Hour(); got != tt.want[3] {
			t.Errorf("hour: want %v got %v for input %v", tt.want[3], got, tt.str)
			failed = true
		}
		if got := parsedTime.Minute(); got != tt.want[4] {
			t.Errorf("minute: want %v got %v for input %v", tt.want[4], got, tt.str)
			failed = true
		}
		if got := parsedTime.Second(); got != tt.want[5] {
			t.Errorf("second: want %v got %v for input %v", tt.want[5], got, tt.str)
			failed = true
		}
		if got := parsedTime.Nanosecond(); got != tt.want[6] {
			t.Errorf("nseconds: want %v got %v for input %v", tt.want[6], got, tt.str)
			failed = true
		}

		convertedTime := parsedTime.In(tt.time_zone)

		if convertedTime.Format(time.RFC3339) != parsedTime.Format(time.RFC3339) {
			t.Errorf("incorrect timezone: want %v got %v for input %v", convertedTime.Location(), parsedTime.Location(), tt.str)
			failed = true
		}

		fmt.Printf("timezone: %v and %v\n", parsedTime, convertedTime)
		
		if failed {
			fmt.Printf("case failed: %v\n", tt.str)
		} else {
			fmt.Printf("case passed: %v\n", tt.str)
		}
	}
}

func TestDatetimeParse(t *testing.T) {

}
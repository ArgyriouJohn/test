package main

import (
	"io/ioutil"
	"time"
)

func formatTime(date string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	parseDate, err := time.Parse(date, layout)

	if err != nil {
		return time.Now()
	} else {
		return parseDate
	}
}

func now() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}

func loadConfig() string {
	b, err := ioutil.ReadFile("config.toml")

	if err != nil {
		Error.Fatalln("Could not load configuration")
	}

	return string(b)
}

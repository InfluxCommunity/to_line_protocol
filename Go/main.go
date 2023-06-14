package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/InfluxCommunity/influxdb3-go/influx"
)

type Value map[string]interface{}

var lines []string // Array to store lines in influx line protocol format

func main() {
	// Use env variables to initialize client
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")
	fmt.Printf("url", url)

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influx.New(influx.Configs{
		HostURL:   url,
		AuthToken: token,
	})

	if err != nil {
		panic(err)
	}
	// Close client at the end and escalate error if present
	defer func(client *influx.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	// Prepare FlightSQL query
	query := `
	    SELECT *
	    FROM "cpu"
	  `

	iterator, err := client.Query(context.Background(), database, query)

	if err != nil {
		panic(err)
	}

	// Specify your configuration here
	measurement := "cpu_test"
	timestamp := "time"
	tags := []string{"host", "cpu"}
	fields := []string{"usage_user"}

	for iterator.Next() {
		value := iterator.Value()

		// Collect tag set
		var tagSet []string
		for _, tag := range tags {
			tagSet = append(tagSet, fmt.Sprintf("%s=%v", tag, value[tag]))
		}

		// Collect field set
		var fieldSet []string
		for _, field := range fields {
			fieldSet = append(fieldSet, fmt.Sprintf("%s=%v", field, value[field]))
		}

		// Here we convert each value into InfluxDB line protocol format.
		line := fmt.Sprintf("%s,%s %s %v",
			measurement,
			strings.Join(tagSet, ","),
			strings.Join(fieldSet, ","),
			value[timestamp])

		fmt.Println(line)
		lines = append(lines, line) // Append the line to lines array
	}

	// Join all lines with newline character
	body := []byte(strings.Join(lines, "\n"))

	// Write to InfluxDB
	err = client.Write(context.Background(), database, body)
	if err != nil {
		panic(err)
	}

}

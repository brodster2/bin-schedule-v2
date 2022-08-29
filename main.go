package main

import (
	"encoding/csv"
	"os"
	"strings"
)

type ScheduleDataSource interface {
	Load(string) ([]Collection, error)
	GetBins(string) (error, Collection)
}

type Collection struct {
	Date string
	Bins []string
}

type CSVDataSource struct{}

func (c CSVDataSource) Load(path string) ([]Collection, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	_, err = reader.Read() // discard column headers
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	//fmt.Println(data)
	var schedule []Collection
	for _, d := range data {
		//fmt.Printf("%d:\t%s\n", k, d)
		var stripped []string
		for _, bin := range strings.Split(d[1], ":") {
			stripped = append(stripped, strings.TrimSpace(bin))
		}
		schedule = append(schedule, Collection{
			Date: d[0],
			Bins: stripped,
		})
	}
	return schedule, nil
}

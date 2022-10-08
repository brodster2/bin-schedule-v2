package helpers

import (
	"encoding/csv"
	"strings"
)

func LoadFromCSV(content string) (map[string][]string, error) {

	reader := csv.NewReader(strings.NewReader(content))

	if _, err := reader.Read(); err != nil { // Discard column header
		return nil, err
	}

	csvContent, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	schedule := make(map[string][]string)

	for _, v := range csvContent {
		var strip []string
		for _, s := range strings.Split(v[1], ":") {
			strip = append(strip, strings.TrimSpace(s))
		}
		schedule[v[0]] = strip
	}

	return schedule, nil
}

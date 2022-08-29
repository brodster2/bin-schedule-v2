package main

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func LoadFromCSV(filepath string) (map[string][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil { // Discard column header
		return nil, err
	}

	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	schedule := make(map[string][]string)

	for _, v := range content {
		var strip []string
		for _, s := range strings.Split(v[1], ":") {
			strip = append(strip, strings.TrimSpace(s))
		}
		schedule[v[0]] = strip
	}

	return schedule, nil
}

var schedule, err = LoadFromCSV("test_schedule.csv")

func main() {
	r := gin.Default()
	r.GET("/bins/:date", func(c *gin.Context) {
		date := c.Param("date")
		bins := schedule[date]
		c.JSON(http.StatusOK, gin.H{
			"bins": bins,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

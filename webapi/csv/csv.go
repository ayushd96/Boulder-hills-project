package csv

import (
	"encoding/csv"
	"os"
	"webapi/trail"
)

func LoadDataFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var trails []trail.Trail
	for i, record := range records {
		if i == 0 {
			continue
		}
		trail := trail.Trail{
			ID:         record[6],
			Name:       record[30],
			BikeTrail:  record[11],
			Fishing:    record[3],
			Difficulty: record[21],
			Fee:        record[9],
			Picnic:     record[1],
			Restrooms:  record[2],
			Address:    record[8],
		}
		trails = append(trails, trail)
	}
	trail.Trails = trails
	return nil
}

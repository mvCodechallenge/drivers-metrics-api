package data_access

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"database/sql"
)

/*
	Holds DB model for Metric data
*/
type DbMetric struct {
	Name      string
	Longitude float64
	Latitude  float64
	Timestamp uint64
	DriverId  uint
	Value     float64
}

/*
	Parse metric data row got DB to DB metric model
	rowData looks like (0,132726768,12.111000000000001,-87.221999999999994,1352)
*/
func parseMetricRow(rowData string) (*DbMetric, error) {
	// Check for size
	if (len(rowData) < 4) {
		return nil, errors.New(fmt.Sprintf("Can't parse metric row data: %s", rowData))
	}

	// Strip delimiters
	dbData := strings.Replace(rowData, "(", "", -1)
	dbData = strings.Replace(dbData, ")", "", -1)
	dbData = strings.Replace(dbData, "\"", "", -1)

	// Check that all metric data parts exists
	if (len(rowData) < 4) {
		return nil, errors.New(fmt.Sprintf("Can't parse metric row data: %s", rowData))
	}

	rowParts := strings.Split(dbData, ",")
	if (len(rowParts) < 5) {
		return nil, errors.New(fmt.Sprintf("Can't parse metric row data: %s", rowData))
	}

	dbm := DbMetric{}

	val, err := strconv.ParseFloat(rowParts[0], 64)
	if (err != nil) {
		return nil, err;
	} else {
		dbm.Value = val
	}

	dbm.Timestamp, err = strconv.ParseUint(rowParts[1], 10, 64)
	if (err != nil) {
		return nil, err;
	}

	dbm.Longitude, err = strconv.ParseFloat(rowParts[2], 64)
	if (err != nil) {
		return nil, err;
	}

	dbm.Latitude, err = strconv.ParseFloat(rowParts[3], 64)
	if (err != nil) {
		return nil, err;
	}

	dId, err := strconv.ParseUint(rowParts[4], 10, 64)
	if (err != nil) {
		return nil, err;
	} else {
		dbm.DriverId = uint(dId)
	}

	return &dbm, nil
}

/*
	Private function to handle iteration and fetching of metrics from DB into array of metric model
*/
func handleMetricsResponse(name string, rows *sql.Rows, err error) ([]*DbMetric, error) {
	if (err != nil) {
		return nil, err;
	} else {
		metrics := []*DbMetric{}
		defer rows.Close()
		for rows.Next() {
			var mbData string;
			err := rows.Scan(&mbData)
			if (err != nil) {
				return nil, err
			} else {
				dbM, err := parseMetricRow(mbData);
				if (err != nil) {
					return nil, err
				} else {
					dbM.Name = name
					metrics = append(metrics, dbM)
				}
			}
		}

		return metrics, nil
	}
}

/*
	Gets metrics data from DB by name of metric and optional driver id (0 means none)
*/
func GetMetrics(name string, id uint) ([]*DbMetric, error) {
	rows, err := _db.Query("select getMetrics($1, $2)", name, id)
	return handleMetricsResponse(name, rows, err);
}

/*
	Gets metrics stat data from DB by name of metric and optional driver id (0 means none)
	and optional stats type (Currently min or max)
*/
func GetMetricsStats(name string, id uint, stats string) ([]*DbMetric, error) {
	rows, err := _db.Query("select getMetricsStats($1, $2, $3)", name, id, stats)
	return handleMetricsResponse(name, rows, err);
}

/*
	Adds new metric data for driver to DB
*/
func AddMetricForDriver(metric DbMetric) error {
	result, err := _db.Query("select insertMetric($1, $2, $3, $4, $5, $6)", metric.DriverId, metric.Name, metric.Timestamp, metric.Longitude, metric.Latitude, metric.Value)
	if (err != nil) {
		return err
	} else {
		result.Close();
		return nil
	}
}
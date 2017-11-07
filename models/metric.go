package models

import (
	"fmt"
	"drivers-metrics-api/data-access"
	"errors"
)

/*
	Holds metric app model
*/
type Metric struct {
	Name        string
	Longitude   float64
	Latitude    float64
	Timestamp   uint64
	DriverId    uint
	Value       float64
}

/*
	Use to print nicely the driver app model info
*/
func (m *Metric) ToString() string {
	return fmt.Sprintf("Name: %s, DriverId: %d, Timestamp: %d, Longitude: %f, Latitude: %f, Value: %f", m.Name, m.DriverId, m.Timestamp, m.Longitude, m.Latitude, m.Value);
}

/*
	Method for construct metric DB model by metric app model
*/
func (m *Metric) CopyToDbModel() data_access.DbMetric {
	return data_access.DbMetric{Name:m.Name, Longitude:m.Longitude, Latitude:m.Latitude, Timestamp:m.Timestamp, DriverId:m.DriverId, Value:m.Value}
}

/*
	Construct metric app model by metric DB model
*/
func copyMetricFromDbModel(dm *data_access.DbMetric) *Metric {
	return &Metric{Name:dm.Name, Longitude:dm.Longitude, Latitude:dm.Latitude, Timestamp:dm.Timestamp, DriverId:dm.DriverId, Value:dm.Value}
}

/*
	Gets metrics info by name and optional drive id (0 means none)
*/
func GetMetrics(name string, id uint) ([]*Metric, error) {
	mDbMetrics, err := data_access.GetMetrics(name, id);
	if (err == nil) {
		metrics := []*Metric{}
		for _, dbM := range mDbMetrics {
			metrics = append(metrics, copyMetricFromDbModel(dbM))
		}

		return metrics, nil
	}

	return nil, errors.New(fmt.Sprintf("Couldn't get metric %s values for driver: %d, extended information: %s", name, id, getInnerError(err)))
}

/*
	Gets metric stats by metric name and optional driver id and stats type (min or max)
*/
func GetMetricsStats(name string, id uint, stats string) ([]*Metric, error) {
	mDbMetrics, err := data_access.GetMetricsStats(name, id, stats);
	if (err == nil) {
		metrics := []*Metric{}
		for _, dbM := range mDbMetrics {
			metrics = append(metrics, copyMetricFromDbModel(dbM))
		}

		return metrics, nil
	}

	return nil, errors.New(fmt.Sprintf("Couldn't get metric: %s with stats: %s for driver: %d, extended information: %s", name, stats, id, getInnerError(err)))
}

/*
	Add new metric data for a driver
*/
func AddMetricForDriver(metric Metric) (error) {
	err := data_access.AddMetricForDriver(metric.CopyToDbModel());
	if (err != nil) {
		return errors.New(fmt.Sprintf("Couldn't add metric: %s, extended information: %s", metric.ToString(), getInnerError(err)))
	}

	return nil
}

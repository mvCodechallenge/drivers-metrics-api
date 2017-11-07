package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"drivers-metrics-api/models"
	"strings"
)

const METRIC_NAME string = ":name";

// Operations about Metrics
type MetricController struct {
	beego.Controller
}

/**
	Get driver id information from request query string parameter "driverId=XXX"
 */
func (m *MetricController) getDriverId() uint {
	var uid uint;
	m.Ctx.Input.Bind(&uid, "driverId");
	return uid;
}

/**
	Get stats type from request query string parameter "stats=min|max"
 */
func (m *MetricController) getStatsType() string {
	var stats string;
	m.Ctx.Input.Bind(&stats, "stats");
	return strings.ToLower(stats);
}

// @Title AddMetricForDriver
// Sample POST /v1/metric/gps.location_lost?driverId=2
// @Description create Metrics for driver
// @Param	body		body 	models.Metric	true		"body for metric content"
// @Success 200 {int} models.Metric.Name
// @Failure 403 body is empty
// @router / [post]
func (m *MetricController) Post() {
	response := APIResponse{}
	name := m.GetString(METRIC_NAME)
	id := m.getDriverId()
	if (id < 1) {
		response.Error = "driverId input must be provided."
	} else {
		var metric models.Metric
		err := json.Unmarshal(m.Ctx.Input.RequestBody, &metric)
		if (err != nil) {
			response.Error = err.Error()
		} else {
			metric.Name = name;
			metric.DriverId = uint(id)
			err := models.AddMetricForDriver(metric)
			if (err != nil) {
				response.Error = err.Error()
			} else {
				response.Data = metric
			}
		}
	}

	m.Data[JSON_DATA] = response;
	m.ServeJSON()
}

// @Title Get
// @Description get metric by name
// Sample 1: /v1/metric/gps.location_lost?driverId=5
// Sample 2: /v1/metric/gps.location_lost
// Sample 3: /v1/metric/gps.location_lost?driverId=5&stats=max
// Sample 4: /v1/metric/gps.location_lost?stats=min
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Metric
// @Failure 403 :name is empty
// @router /:name [get]
func (m *MetricController) Get() {
	response := APIResponse{}
	name := m.GetString(METRIC_NAME)
	stats := m.getStatsType()

	var metrics []*models.Metric
	var err error

	/*
		If Stats parameter exists get stats call (min or max for now) on metric name (and on driver id if exists),
		if not do a get a metrics get on metric name (and on driver id if exists),
	*/
	if (stats != "") {
		metrics, err = models.GetMetricsStats(name, m.getDriverId(), stats)
	} else {
		metrics, err = models.GetMetrics(name, m.getDriverId())
	}

	if (err != nil) {
		response.Error = err.Error()
	} else {
		response.Data = metrics
	}


	m.Data[JSON_DATA] = response
	m.ServeJSON()
}
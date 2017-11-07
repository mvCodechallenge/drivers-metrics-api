package controllers

import (
	"CodeChallenge/DriversMetricsAPI/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

const DRIVER_ID string = ":id";

// Operations about Drivers
type DriverController struct {
	beego.Controller
}

// @Title CreateDriver
// @Description create Drivers
// @Param	body		body 	models.Driver	true		"body for driver content"
// @Success 200 {int} models.Driver.Id
// @Failure 403 body is empty
// @router / [post]
func (d *DriverController) Post() {
	response := APIResponse{}
	var driver models.Driver
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &driver)
	if (err != nil) {
		response.Error = err.Error()
	} else {
		id, err := models.AddDriver(driver)
		if (err != nil) {
			response.Error = err.Error()
		} else {
			driver.Id = id
			response.Data = driver
		}
	}

	d.Data[JSON_DATA] = response;
	d.ServeJSON()
}

// @Title GetAllDrivers
// @Description get all Drivers
// @Success 200 {object} models.Driver
// @router / [get]
func (d *DriverController) GetAllDrivers() {
	response := APIResponse{}
	drivers, err := models.GetAllDrivers()
	if (err != nil) {
		response.Error = err.Error()
	} else {
		response.Data = drivers
	}

	d.Data[JSON_DATA] = response;
	d.ServeJSON()
}

// @Title Get
// @Description get driver by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Driver
// @Failure 403 :id is empty
// @router /:id [get]
func (d *DriverController) Get() {
	response := APIResponse{}
	id, err := d.GetUint32(DRIVER_ID)
	if (err != nil) {
		response.Error = err.Error()
	} else {
		driver, err := models.GetDriver(uint(id))
		if (err != nil) {
			response.Error = err.Error()
		} else {
			response.Data = driver
		}
	}

	d.Data[JSON_DATA] = response
	d.ServeJSON()
}

// @Title Update
// @Description update the driver
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Driver	true		"body for driver content"
// @Success 200 {object} models.Driver
// @Failure 403 :id is not uint
// @router /:id [put]
func (d *DriverController) Put() {
	response := APIResponse{}
	id, err := d.GetUint32(DRIVER_ID)
	if (err != nil) {
		response.Error = err.Error()
	} else {
		var driver models.Driver
		err = json.Unmarshal(d.Ctx.Input.RequestBody, &driver)
		if (err != nil) {
			response.Error = err.Error()
		} else {
			driver.Id = uint(id)
			err := models.UpdateDriver(driver)
			if (err != nil) {
				response.Error = err.Error()
			} else {
				// If OK updating driver then try to make a GET request on it
				d.Get()
				return
			}
		}
	}

	d.Data[JSON_DATA] = response
	d.ServeJSON()
}

// @Title Delete
// @Description delete the driver
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (d *DriverController) Delete() {
	response := APIResponse{}
	id, err := d.GetUint32(DRIVER_ID)
	if (err != nil) {
		response.Error = err.Error()
	} else {
		err := models.DeleteDriver(uint(id))
		if (err != nil) {
			response.Error = err.Error()
		} else {
			// If OK deleting driver then try to make a GET ALL request on it to see that it's missing
			d.GetAllDrivers()
			return
		}
	}

	d.Data[JSON_DATA] = response
	d.ServeJSON()
}



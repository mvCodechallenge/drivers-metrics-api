package models

import (
	"fmt"
	"errors"
	"drivers-metrics-api/data-access"
)

/*
	Holds Driver app model
*/
type Driver struct {
	Id       	    uint
	Name     	    string
	LicenseNumber	string
}

/*
	Use to print nicely the driver app model info
*/
func (d *Driver) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, LicenseNumber: %s", d.Id, d.Name, d.LicenseNumber);
}

/*
	Method for construct driver DB model by driver app model
*/
func (d *Driver) CopyToDbModel() data_access.DbDriver {
	return data_access.DbDriver{Id:d.Id, Name:d.Name, LicenseNumber:d.LicenseNumber}
}

/*
	Construct driver app model by driver DB model
*/
func copyDriverFromDbModel(dd *data_access.DbDriver) *Driver {
	return &Driver{Id:dd.Id, Name:dd.Name, LicenseNumber:dd.LicenseNumber}
}

/*
	Get driver info by it's id
*/
func GetDriver(id uint) (*Driver, error) {
	dd, err := data_access.GetDriver(id);
	if (err == nil) {
		return copyDriverFromDbModel(dd), nil
	}



	return nil, errors.New(fmt.Sprintf("Driver %d doesn't exists, extended information: %s", id, getInnerError(err)))
}

/*
	Gets all drivers info
*/
func GetAllDrivers() (map[uint]*Driver, error) {
	dbDrivers, err := data_access.GetAllDrivers()
	if (err != nil) {
		fmt.Printf(err.Error());
		err = errors.New(fmt.Sprintf("Couldn't get all drivers, extended information: %s", getInnerError(err)))
	}

	mapOfDrivers := make(map[uint]*Driver)
	for key, val := range dbDrivers{
		mapOfDrivers[key] = copyDriverFromDbModel(val)
	}


	return mapOfDrivers, err
}

/*
	Adds new driver
*/
func AddDriver(d Driver) (uint, error) {
	id, err := data_access.AddDriver(d.CopyToDbModel());
	if (err != nil) {
		return 0, errors.New(fmt.Sprintf("Couldn't add driver: %s, extended information: %s", d.ToString(), getInnerError(err)))
	}

	return id, nil
}

/*
	Updates existing driver info
*/
func UpdateDriver(d Driver) (error) {
	err := data_access.UpdateDriver(d.CopyToDbModel());
	if (err != nil) {
		fmt.Printf(err.Error());
		err = errors.New(fmt.Sprintf("Couldn't update driver: %d, extended information: %s", d.Id, getInnerError(err)))
	}

	return err
}

/*
	Deletes driver by id
*/
func DeleteDriver(id uint) error {
	err := data_access.DeleteDriver(id);
	if (err != nil) {
		fmt.Printf(err.Error());
		err = errors.New(fmt.Sprintf("Couldn't delete driver: %d, extended information: %s", id, getInnerError(err)))
	}

	return err
}

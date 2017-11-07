package data_access

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
)

/*
	Holds DB model for Driver
*/
type DbDriver struct {
	Id              uint
	Name            string
	LicenseNumber   string
}

/*
	Parse driver row data got DB to DB Driver model
	rowData looks like (1366,"New One",12-288-10)
*/
func parseDriverRow(rowData string) (*DbDriver, error) {
	// Check for size
	if (len(rowData) < 3) {
		return nil, errors.New(fmt.Sprintf("Can't parse driver row data: %s", rowData))
	}

	// Strip delimiters
	dbData := strings.Replace(rowData, "(", "", -1)
	dbData = strings.Replace(dbData, ")", "", -1)
	dbData = strings.Replace(dbData, "\"", "", -1)

	// Check that it has all driver info parts
	if (len(rowData) < 3) {
		return nil, errors.New(fmt.Sprintf("Can't parse driver row data: %s", rowData))
	}

	rowParts := strings.Split(dbData, ",")
	if (len(rowParts) < 3) {
		return nil, errors.New(fmt.Sprintf("Can't parse driver row data: %s", rowData))
	}

	dbd := DbDriver{}
	id, _ := strconv.ParseUint(rowParts[0], 10, 64)
	dbd.Id = uint(id)
	dbd.Name = rowParts[1]
	dbd.LicenseNumber = rowParts[2]

	return &dbd, nil
}

/*
	Get driver info from DB
*/
func GetDriver(id uint) (*DbDriver, error) {
	row := _db.QueryRow("select getDriverById($1)", id)
	var dbData string;
	err := row.Scan(&dbData)
	if (err != nil) {
		return nil, err
	} else {
		driver, err := parseDriverRow(dbData)
		if (err != nil) {
			return nil, err
		} else {
			return driver, nil
		}
	}
}

/*
	Get all drivers info from DB
*/
func GetAllDrivers() (map[uint]*DbDriver, error) {
	rows, err := _db.Query("select getAllDrivers()")
	if (err != nil) {
		return nil, err;
	} else {
		drivers := make(map[uint]*DbDriver)
		defer rows.Close()
		for rows.Next() {
			var dbData string;
			err := rows.Scan(&dbData)
			if (err != nil) {
				return nil, err
			} else {
				driver, err := parseDriverRow(dbData)
				if (err != nil) {
					return nil, err
				} else {
					drivers[driver.Id] = driver
				}
			}
		}

		return drivers, nil
	}
}

/*
	Add new driver info to DB
*/
func AddDriver(d DbDriver) (uint, error) {
	result, err := _db.Query("select addDriver($1, $2)", d.Name, d.LicenseNumber)
	if (err != nil) {
		return 0, err;
	} else {
		defer result.Close();
		for result.Next() {
			var id uint;
			err = result.Scan(&id);
			if (err != nil) {
				return 0, err
			}

			return id, nil
		}
	}

	return 0, errors.New("Couldn't add driver to storage");
}

/*
	Update existing driver info to DB
*/
func UpdateDriver(d DbDriver) (error) {
	_, err := _db.Exec("select updateDriver($1, $2, $3)", d.Id, d.Name, d.LicenseNumber)
	return err
}

/*
	Delete driver info from DB
*/
func DeleteDriver(id uint) error {
	_, err := _db.Exec("select deleteDriver($1)", id)
	return err
}


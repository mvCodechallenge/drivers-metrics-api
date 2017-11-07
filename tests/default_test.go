package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "CodeChallenge/DriversMetricsAPI/routers"
	_ "CodeChallenge/DriversMetricsAPI/data-access"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"bytes"
	"encoding/json"
	"CodeChallenge/DriversMetricsAPI/models"
	"CodeChallenge/DriversMetricsAPI/controllers"
	"github.com/smartystreets/assertions/should"
	"fmt"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func parseDriver(res interface{}) models.Driver {
	fmt.Printf("----- res type is [%T] and value [%v] -----", res, res);
	var driverObj  = res.(map[string]interface{})
	fmt.Printf("----- driverObj type is [%T] and value [%v] -----", driverObj, driverObj);
	return models.Driver{Id: uint(driverObj["Id"].(float64)), Name: driverObj["Name"].(string), LicenseNumber: driverObj["LicenseNumber"].(string)}
}

// TestGet is a sample to run an endpoint test
func TestGetAll(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/driver", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetAll", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	})
}

func TestPost(t *testing.T) {
	r, _ := http.NewRequest("POST", "/v1/driver/", bytes.NewReader([]byte("{\"Name\": \"New One\", \"LicenseNumber\": \"12-288-10\"}")))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

		Convey("The Result Should be new driver id", func() {
			res := controllers.APIResponse{}
			json.Unmarshal(w.Body.Bytes(), &res)
			driver := parseDriver(res.Data)

			So(driver.Id, should.BeGreaterThan, 0)
			fmt.Printf(" *** Added new driver id: %d", driver.Id)

			So(driver.Name, ShouldEqual, "New One")
			So(driver.LicenseNumber, ShouldEqual, "12-288-10")

		})
	})
}

func TestPut(t *testing.T) {
    url := "/v1/driver/2"
	fmt.Printf(" *** URL on PUT: %s", url)

	r, _ := http.NewRequest("PUT", url, bytes.NewReader([]byte("{\"Name\": \"Updated Name\", \"LicenseNumber\": \"12-999-10\"}")))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

		Convey("The Result Should be what we changed", func() {
			res := controllers.APIResponse{}
			json.Unmarshal(w.Body.Bytes(), &res)
			driver := parseDriver(res.Data)

			So(driver.Id, ShouldEqual, 2)
			So(driver.Name, ShouldEqual, "Updated Name")
			So(driver.LicenseNumber, ShouldEqual, "12-999-10")
		})
	})
}

func TestGet(t *testing.T) {
	url := "/v1/driver/2"
	fmt.Printf(" *** URL on GET: %s", url)

	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should be what we get", func() {
			res := controllers.APIResponse{}
			json.Unmarshal(w.Body.Bytes(), &res)
			driver := parseDriver(res.Data)

			So(driver.Id, ShouldEqual, 2)
			So(driver.Name, ShouldEqual, "Updated Name")
			So(driver.LicenseNumber, ShouldEqual, "12-999-10")
		})
	})
}

func TestDelete(t *testing.T) {
	url := "/v1/driver/9"
	fmt.Printf(" *** URL on GET: %s", url)

	r, _ := http.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}



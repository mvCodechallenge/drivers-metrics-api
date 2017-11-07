package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"],
		beego.ControllerComments{
			Method: "GetAllDrivers",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:DriverController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:MetricController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:MetricController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["drivers-metrics-api/controllers:MetricController"] = append(beego.GlobalControllerRouter["drivers-metrics-api/controllers:MetricController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

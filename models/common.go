package models

import "github.com/astaxie/beego"

/*
	When showing errors on Response we would like to return also the real error when on dev run-mode
*/
func getInnerError(err error) string {
	innerError := "None"
	if (beego.BConfig.RunMode == "dev") {
		innerError = err.Error()
	}

	return innerError
}

// @APIVersion 1.0.0
// @Title CoParent API
// @Description API for CoParent app
// @Contact devaris@coparent.io
// @TermsOfServiceUrl http://coparent.io/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"os"

	"github.com/devarispbrown/coparent-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "" {
		runMode = "dev"
	}
	beego.BConfig.RunMode = runMode

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/parents",
			beego.NSInclude(
				&controllers.ParentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"],
		beego.ControllerComments{
			Method: "GetAllParents",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"],
		beego.ControllerComments{
			Method: "CreateParent",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"],
		beego.ControllerComments{
			Method: "GetOneParent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"],
		beego.ControllerComments{
			Method: "UpdateParent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"] = append(beego.GlobalControllerRouter["github.com/devarispbrown/coparent-api/controllers:ParentController"],
		beego.ControllerComments{
			Method: "DeleteParent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}

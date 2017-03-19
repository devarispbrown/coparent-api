package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/devarispbrown/coparent-api/models"
)

type ParentController struct {
	beego.Controller
}

// @router / [get]
func (o *ParentController) GetAllParents() {
	obs := models.GetAllParents()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @router / [post]
func (o *ParentController) CreateParent() {
	var parent models.Parent
	o.ParseForm(&parent)
	result, err := models.CreateParent(parent)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// @router /:id [get]
func (o *ParentController) GetOneParent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	user, err := models.GetOneParent(uid)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = user
	}

	o.ServeJSON()
}

// @router /:id [put]
func (o *ParentController) UpdateParent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}
	id, _ := strconv.Atoi(uid)
	parent := models.Parent{Id: id}

	o.ParseForm(&parent)
	parent, err := models.UpdateParent(parent)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = parent
	}
}

// @router /:id [delete]
func (o *ParentController) DeleteParent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	id, error := strconv.Atoi(uid)
	if error != nil {
		o.Abort("500")
	}
	response, error := models.DeleteParent(id)
	if error == nil {
		o.Data["json"] = response
	} else {
		o.Abort("500")
	}

	o.ServeJSON()
}

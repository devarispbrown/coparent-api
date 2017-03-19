package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Parent struct {
	Id            int `orm:"index"`
	FirstName     string
	LastName      string
	Gender        string
	TotalIncome   int
	PhoneNumber   string
	EmailAddress  string
	StreetAddress string
	City          string
	State         string
	ZipCode       string
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Parent))
}

func (u *Parent) TableName() string {
	return "parents"
}

func GetAllParents() []*Parent {
	var parents []*Parent
	o := orm.NewOrm()
	qs := o.QueryTable("parents")
	num, err := qs.All(&parents)
	if num > 0 {
		fmt.Println(err)
	}

	return parents
}

func CreateParent(parent Parent) (p Parent, err error) {
	o := orm.NewOrm()
	fmt.Println(parent)

	id, err := o.Insert(&parent)
	if err == nil {
		fmt.Println(id)
	}

	return parent, nil

}

func GetOneParent(id string) (Parent, error) {
	uid, _ := strconv.Atoi(id)
	parent := Parent{Id: uid}
	o := orm.NewOrm()
	err := o.Read(&parent)
	if err == orm.ErrNoRows {
		fmt.Println(errors.New("not"))
		return parent, errors.New("404")
	} else {
		return parent, nil
	}
}

func UpdateParent(parent Parent) (p Parent, err error) {
	fmt.Println(parent.Id)
	o := orm.NewOrm()
	id, err := o.Update(&parent)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
	return parent, err
}

func DeleteParent(id int) (bool, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Parent{Id: id}); err == nil {
		fmt.Println(num)
		return true, nil
	}
	return false, errors.New("Error")
}

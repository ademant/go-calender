package models

import (
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
)

//a struct to rep user account
type Permission struct {
	gorm.Model
	ShortCode string `json:"ShortCode" gorm:"unique"`
	Description string `json:"Description"`
}

func (perm *Permission) DBPermissionInit(){
	// initial set of permissions
	UserAdd := Permission{ShortCode:"AddUser", Description: "Add or Update User"}
	UserDel := Permission{ShortCode:"DelUser", Description: "Delete User"}
	UserGet := Permission{ShortCode:"GetUser", Description: "Read User informations"}
	GetDB().Where(Permission{ShortCode:"AddUser"}).FirstOrCreate(&UserAdd)
	GetDB().Where(Permission{ShortCode:"DelUser"}).FirstOrCreate(&UserDel)
	GetDB().Where(Permission{ShortCode:"GetUser"}).FirstOrCreate(&UserGet)
}
//Validate incoming user details...
func (perm *Permission) Validate() (map[string]interface{}, bool) {

	//Role name must be unique
	temp := &Permission{}

	//check for errors and duplicate roles
	err := GetDB().Table("permissions").Where("ShortCode = ?", perm.ShortCode).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Description != "" {
		return u.Message(false, "Permission description."), false
	}

	return u.Message(false, "Requirement for permission passed"), true
}

func (perm *Permission) Create() (map[string]interface{}) {

	if resp, ok := perm.Validate(); !ok {
		return resp
	}

	GetDB().Create(perm)

	if perm.ID <= 0 {
		return u.Message(false, "Failed to create permission, connection error.")
	}

	response := u.Message(true, "Permission has been created")
	response["permission"] = perm
	return response
}

func GetPermission(u uint) *Permission {
	perm := &Permission{}
	GetDB().Table("permissions").Where("id = ?", u).First(perm)
	if perm.ShortCode == "" { //Permission not found!
		return nil
	}
	return perm
}

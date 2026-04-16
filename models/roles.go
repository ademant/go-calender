package models

import (
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
)

//a struct to rep user account
type Roles struct {
	gorm.Model
	Tier uint `json:"Level" gorm:"unique"`
	Role string `json:"Role"`
}

//Validate incoming user details...
func (role *Roles) Validate() (map[string]interface{}, bool) {

	//Role name must be unique
	temp := &Roles{}

	//check for errors and duplicate roles
	err := GetDB().Table("roles").Where("role = ?", role.Role).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Role != "" {
		return u.Message(false, "Role description already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (role *Roles) Create() (map[string]interface{}) {

	if resp, ok := role.Validate(); !ok {
		return resp
	}

	GetDB().Create(role)

	if role.ID <= 0 {
		return u.Message(false, "Failed to create role, connection error.")
	}

	response := u.Message(true, "Role has been created")
	response["role"] = role
	return response
}

func GetRole(u uint) *Roles {
	role := &Roles{}
	GetDB().Table("roles").Where("id = ?", u).First(role)
	if role.Role == "" { //User not found!
		return nil
	}

	return role
}

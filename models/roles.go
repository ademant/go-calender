package models

import (
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
)

//a struct to rep user account
type Role struct {
	gorm.Model
	Tier uint `json:"Level" gorm:"unique"`
	Role string `json:"Role"`
}

func (role *Role) DBRoleInit() {
// Ensure admin role exist
	initRole := Role{Tier:0, Role:"admin"}
	GetDB().Where(Role{Tier:0}).FirstOrCreate(&initRole)
}

//Validate incoming user details...
func (role *Role) Validate() (map[string]interface{}, bool) {

	//Role name must be unique
	temp := &Role{}

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

func (role *Role) Create() (map[string]interface{}) {

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

func GetRole(u uint) *Role {
	role := &Role{}
	GetDB().Table("roles").Where("id = ?", u).First(role)
	if role.Role == "" { //User not found!
		return nil
	}

	return role
}

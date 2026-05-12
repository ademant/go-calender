package models

import (
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
)

//a struct to rep user account
type AccountRoles struct {
	gorm.Model
	RoleID int `json:"roleid"`
	AccountID int `json:"accountid" gorm:"unique"`
}

//Validate incoming user details...
func (ar *AccountRoles) Validate() (map[string]interface{}, bool) {

	//Role name must be unique
	temp := &AccountRoles{}

	//check for errors and duplicate roles
	err := GetDB().Table("accountroles").Where("accountid = ?", ar.AccountID).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if err == nil  {
		return u.Message(false, "Role description already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (ar *AccountRoles) Create() (map[string]interface{}) {

	if resp, ok := ar.Validate(); !ok {
		return resp
	}

	GetDB().Create(ar)

	if ar.ID <= 0 {
		return u.Message(false, "Failed to create role, connection error.")
	}

	response := u.Message(true, "Role has been created")
	response["role"] = ar
	return response
}

func GetAccountRole(u uint) *AccountRoles {
	ar := &AccountRoles{}
	err := GetDB().Table("accountroles").Where("id = ?", u).First(ar)
	if err != nil { 
		return nil
	}

	return ar
}

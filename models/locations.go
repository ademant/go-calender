package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
)

type Location struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique"`
	Address  string `json:"address"`
	Room string `json:"room"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (location *Location) Validate() (map[string]interface{}, bool) {

	if location.Name == "" {
		return u.Message(false, "Location name should be on the payload"), false
	}

	if location.Address == "" {
		return u.Message(false, "Location address should be on the payload"), false
	}

	if location.Latitude > 90 && location.Latitude < -90 {
		return u.Message(false, "Latitude is out of range"), false
	}

	if location.Longitude > 180 && location.Longitude < -180 {
		return u.Message(false, "Longitude is out of range"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (location *Location) Create() (map[string]interface{}) {

	if resp, ok := location.Validate(); !ok {
		return resp
	}

	dbresp := GetDB().Create(location)
	if dbresp.Error != nil {
		fmt.Print(dbresp.Error)
		return u.Message(false,fmt.Sprintf("%v",dbresp.Error))
	}
	fmt.Print("Create Location ID: %s", location.ID)
	resp := u.Message(true, "success")
	resp["location"] = location
	return resp
}

func (location *Location) Update() (map[string]interface{}) {
	if resp, ok := location.Validate(); !ok {
		return resp
	}
	fmt.Print("Update Location ID: %s", location.ID)
	dbresp := db.Save(&location)
	if dbresp.Error != nil {
		return u.Message(false,"Could not update Location")
	}
	resp := u.Message(true, "success")
	resp["location"] = location
	return resp
}

func (location *Location) Delete() (map[string]interface{}) {
	if resp, ok := location.Validate(); !ok {
		return resp
	}
	fmt.Print("Delete Location ID: %s", location.ID)
	dbresp := db.Delete(&location)
	if dbresp.Error != nil {
		return u.Message(false,"Could not delete Location entry ")
	}
	resp := u.Message(true, "success")
	return resp
}

func GetLocation(id uint) (*Location) {

	location := &Location{}
	err := GetDB().Table("locations").Where("id = ?", id).First(location).Error
	if err != nil {
		return nil
	}
	return location
}

func GetLocations() ([]*Location) {

	locations := make([]*Location, 0)
	err := GetDB().Table("locations").Find(&locations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return locations
}

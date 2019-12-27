package model

import "encoding/json"

type Profile struct {
	Name string
	City string
	Gender string
	Birthplace string
	Age 	string
	Education string
	Marriage string
	Height  string
	Income  string
	HaveHouse string
	HaveCar string
	Url string
}

func FromJsonObj (o interface{}) (Profile, error) {
	var profile Profile
	b,err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(b, &profile)

	if err != nil {
		return profile, err
	}

	return profile, nil
}
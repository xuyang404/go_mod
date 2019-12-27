package parser

import (
	"spider/engine"
)

type CityUserParse struct{
	cityName string
}

func (c *CityUserParse) Parse(content []byte, url string) engine.ParseResult {
	return cityUserParse(content, url, c.cityName)
}

func (c *CityUserParse) Serialize() (name string, Args interface{}) {
	return "CityUserParse", c.cityName
}

func NewCityUserParse(name string) *CityUserParse {
	return &CityUserParse{
		cityName: name,
	}
}


type ProfileParseArgs struct{
	UserName string
	CityName string
	Gender string
}

type ProfileParse struct{
	userName string
	cityName string
	gender string
}

func (p * ProfileParse) Parse(content []byte, url string) engine.ParseResult {
	return profileParse(content, url, ProfileParseArgs{
		UserName: p.userName,
		CityName: p.cityName,
		Gender: p.gender,
	})
}

func (p * ProfileParse) Serialize() (name string, Args interface{}) {
	return "ProfileParse", ProfileParseArgs{
		UserName: p.userName,
		CityName: p.cityName,
		Gender: p.gender,
	}
}

func NewProfileParse(userName string, cityName string, gender string) *ProfileParse {
	return &ProfileParse{
		userName: userName,
		cityName: cityName,
		gender:   gender,
	}
}
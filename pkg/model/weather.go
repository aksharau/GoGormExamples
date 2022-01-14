package model

import (
	"reflect"
	"time"
)

type Weather struct {
	City    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Temp struct {
		Tempr    string `json:"temp"`
		Humidity string `json:"humidity"`
	} `json:"main"`
}

type CityWeather struct {
	Date        string `json:"column:CREATE_DATE"`
	City        string `json:"column:CITY"`
	WeatherDesc string `json:"column:WEATHER"`
	Temp        string `json:"column:TEMPR"`
	Humidity    string `json:"column:HUMIDITY"`
}

func (CityWeather) TableName() string {
	return "CITY_WEATHER"
}

func MapToCityWeather(i Weather) CityWeather {
	c := CityWeather{}
	c.City = i.City
	c.Date = time.Now().Format(time.RFC3339)
	c.Humidity = i.Temp.Humidity
	c.Temp = i.Temp.Tempr
	if i.Weather != nil {
		c.WeatherDesc = i.Weather[0].Description
	}

	return c
}

func (c *CityWeather) GetAllFields() []string {
	s := reflect.ValueOf(c).Elem()
	numFields := s.NumField()
	typeOfT := s.Type()

	fields := make([]string, numFields)

	for i := 0; i < numFields; i++ {

		fields[i] = typeOfT.Field(i).Name
	}

	return fields
}

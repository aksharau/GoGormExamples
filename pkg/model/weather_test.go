package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonParsing(t *testing.T) {
	content, err := ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	var responseObject Weather
	json.Unmarshal(content, &responseObject)
	assert.Equal(t, responseObject.City, "London")

}

func TestModelMapping(t *testing.T) {
	content, err := ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	var responseObject Weather
	json.Unmarshal(content, &responseObject)
	assert.Equal(t, responseObject.City, "London")

	city := MapToCityWeather(responseObject)

	assert.Equal(t, city.City, "London")

}

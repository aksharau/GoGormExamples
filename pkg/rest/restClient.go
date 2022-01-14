package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aksharau/GoGormExamples/pkg/model"
)

type RestClient struct {
	client http.Client
	apiKey string
}

func GetRestClient(akey string) RestClient {
	r := RestClient{}
	r.client = http.Client{}
	r.apiKey = akey
	return r
}

func (r *RestClient) GetWeather(city string) model.Weather {

	reqStr := getURL(city, r.apiKey)
	req, err := http.NewRequest("GET", reqStr, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := r.client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	var responseObject model.Weather
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Println("API resposne struct: ", responseObject)

	return responseObject

}

func getURL(city string, key string) string {

	str := "api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + key

	return str
}

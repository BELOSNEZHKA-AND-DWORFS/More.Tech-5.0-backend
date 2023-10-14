package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"more.tech/structs"
	"more.tech/util"
)

func makeRequest(text string, latitude float32, longitude float32) string {
	url := "https://search-maps.yandex.ru/v1/"
	apiKey := "065efd96-9de0-43ed-bf9d-881c088a856a" // TODO: remove hardcoded api key
	coordinates := fmt.Sprintf("%f,%f", longitude, latitude)
	spn := "0.1,0.1" // (угловое расстояние при поиске) // TODO: don't hardcode
	request := fmt.Sprintf("%s/?apikey=%s&text=%s&lang=ru_RU&ll=%s&spn=%s", url, apiKey, text, coordinates, spn)

	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func prepareData(jsonResponse string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonResponse), &data); err != nil {
		fmt.Println("Error:", err)
	}

	flatMap := make(map[string]interface{})
	util.JsonToFlatMap(data, "", flatMap)

	return flatMap
}

func putOfficeDataInStructure(jsonResponse string) []structs.OfficeYandexData {
	flatMap := prepareData(jsonResponse)

	titles := util.GetFeatures(flatMap, "CompanyMetaData.name")
	addresses := util.GetFeatures(flatMap, "CompanyMetaData.address")
	dateTimes := util.GetFeatures(flatMap, "Hours.text")
	phones := util.GetFeatures(flatMap, "Phones")

	size := min(len(titles), len(addresses), len(dateTimes), len(phones))

	var resultData []structs.OfficeYandexData
	for i := 0; i < size; i++ {
		var yandexData structs.OfficeYandexData
		yandexData.Title = titles[i]
		yandexData.Datetime = dateTimes[i]
		yandexData.Phone = phones[i]

		resultData = append(resultData, yandexData)
	}
	return resultData
}

func getOfficeByTitleInArea(title string, latitude float32, longitude float32) []structs.OfficeYandexData {
	requestText := makeRequest(title, latitude, longitude)
	data := putOfficeDataInStructure(requestText)
	return data
}

func putAtmDataInStructure(jsonResponse string) []structs.AtmYandexData {
	flatMap := prepareData(jsonResponse)

	titles := util.GetFeatures(flatMap, "CompanyMetaData.name")
	addresses := util.GetFeatures(flatMap, "CompanyMetaData.address")
	dateTimes := util.GetFeatures(flatMap, "Hours.text")
	phones := util.GetFeatures(flatMap, "Phones")

	size := min(len(titles), len(addresses), len(dateTimes), len(phones))

	var resultData []structs.AtmYandexData
	for i := 0; i < size; i++ {
		var yandexData structs.AtmYandexData
		yandexData.Title = titles[i]
		yandexData.Datetime = dateTimes[i]
		yandexData.Phone = phones[i]

		resultData = append(resultData, yandexData)
	}
	return resultData
}

func getAtmByTitleInArea(title string, latitude float32, longitude float32) []structs.AtmYandexData {
	requestText := makeRequest(title, latitude, longitude)
	data := putAtmDataInStructure(requestText)
	return data
}

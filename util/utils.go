package util

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func JsonToFlatMap(data interface{}, prefix string, result map[string]interface{}) {
	switch value := data.(type) {
	case map[string]interface{}:
		for k, v := range value {
			key := fmt.Sprintf("%s%s", prefix, k)
			JsonToFlatMap(v, key+".", result)
		}
	case []interface{}:
		for i, v := range value {
			key := fmt.Sprintf("%s%d", prefix, i)
			JsonToFlatMap(v, key+".", result)
		}
	default:
		result[prefix[:len(prefix)-1]] = value
	}
}

func GetFeatures(data map[string]interface{}, featureName string) []string {
	var resultMap map[string]string = make(map[string]string)
	var result []string

	for k, v := range data {
		if strings.Index(k, featureName) < 0 {
			continue
		}
		resultMap[k] = fmt.Sprintf("%v", v)
	}

	var keys []string
	for k := range resultMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		result = append(result, resultMap[k])
	}

	return result
}

func CheckOneFilter(service []string, target string) bool {
	for _, s := range service {
		if target == s {
			return true
		}
	}
	return false
}

func CheckFilter(services []string, target []string) bool {
	for _, t := range target {
		if !CheckOneFilter(services, t) {
			return false
		}
	}
	return true
}

const EarthRadius = 6371 // Радиус Земли в километрах

func degreeToRadian(degree float64) float64 {
	return degree * math.Pi / 180
}

func Distance(loc1 []float64, loc2 []float64) float64 {
	lat1Rad := degreeToRadian(loc1[0])
	lon1Rad := degreeToRadian(loc1[1])
	lat2Rad := degreeToRadian(loc2[0])
	lon2Rad := degreeToRadian(loc2[1])

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadius * c
	return distance
}

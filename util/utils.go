package util

import (
	"fmt"
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

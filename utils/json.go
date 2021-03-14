package utils

import (
	"encoding/json"
)

func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		return nil
	}
	return tempMap
	// fmt.Println(tempMap)
	// fmt.Println(tempMap["ret"])

}

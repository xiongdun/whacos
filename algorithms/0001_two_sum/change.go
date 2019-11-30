package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	stringMap := map[string]int{"A": 1, "B.A": 2, "B.B": 3, "CC.D.E": 4, "CC.D.F.G": 5}

	bytes, err := json.MarshalIndent(stringMap, " ", " ")
	fmt.Println("============== before===============")
	fmt.Println(string(bytes))
	if err != nil {
		return
	}

	resultMap := change2(stringMap)

	bytes, err = json.MarshalIndent(resultMap, "	", "	")

	if err != nil {
		return
	}
	fmt.Println("============== after ===============")
	fmt.Println(string(bytes))
}

func change(stringMap map[string]int) map[string]interface{} {

	if stringMap == nil {
		return nil
	}

	resultMap := make(map[string]interface{})

	for key, value := range stringMap {

		val := value
		keys := strings.Split(key, ".")

		if len(keys) == 1 {
			resultMap[keys[0]] = val
		} else {

			var current map[string]interface{}
			if _, ok := resultMap[keys[0]]; !ok {
				for i := 0; i < len(keys); i++ {
					if i == 0 {
						current = make(map[string]interface{})
					}

					if i != len(keys)-1 {
						current[keys[i]] = make(map[string]interface{})
					} else {
						current[keys[i]] = val
					}
				}
			} else {
				current = resultMap[keys[0]].(map[string]interface{})

				for i := 1; i < len(keys); i++ {
					if _, ok := current[keys[i]]; ok && i != len(keys)-1 {
						current[keys[i]] = current[keys[i]].(map[string]interface{})
					} else {
						current[keys[i]] = val
					}
				}
			}

			resultMap[keys[0]] = current
		}
	}
	return resultMap
}

func change2(stringMap map[string]int) map[string]interface{} {
	if stringMap == nil {
		return nil
	}

	resultMap := make(map[string]interface{})

	for key, value := range stringMap {

		keys := strings.Split(key, ".")
		current := resultMap
		for index, k := range keys {
			val := current[k]
			if val == nil {
				if index == len(keys)-1 {
					current[k] = value
					break
				}
				current[k] = make(map[string]interface{})
			}
			current = current[k].(map[string]interface{})
		}
	}

	return resultMap
}

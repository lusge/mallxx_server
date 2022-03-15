package conversion

import (
	"encoding/json"
	"fmt"
)

func StructToMap(src interface{}) (map[string]interface{}, error) {
	// m.redis.Hset(itemKey)
	data, _ := json.Marshal(src)
	js := make(map[string]interface{})

	err := json.Unmarshal(data, &js)

	if err != nil {
		return nil, err
	}

	return js, nil
}

func StructToMapAndValueIsString(s interface{}) (map[string]string, error) {
	data, err := StructToMap(s)

	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for key, value := range data {
		result[key] = InterfaceToString(value)
	}

	return result, nil
}

//
func InterfaceToString(s interface{}) string {
	switch s.(type) {
	case string:
		return s.(string)
	case int, int64:
		return fmt.Sprintf("%d", s)
	case float32, float64:
		return fmt.Sprintf("%f", s)
	default:
		return ""
	}

}

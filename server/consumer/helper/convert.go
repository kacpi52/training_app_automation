package helper

import (
	"encoding/json"
)

func BindJSONToMap(obj interface{}) (map[string]interface{}, error) {
	marshaledData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(marshaledData, &jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}


func UnmarshalJSONToType(jsonStr string, target interface{}) error {

	err := json.Unmarshal([]byte(jsonStr), target);
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonStr), target)

}


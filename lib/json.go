package lib

import (
	"encoding/json"
	"errors"
	"strings"
)

func ParseJson(jsonString string) (map[string]interface{}, error) {
	if strings.TrimSpace(jsonString) == "" {
		return nil, errors.New("invalid json string")
	}
	jsonData := []byte(jsonString)
	var v interface{}
	err := json.Unmarshal(jsonData, &v)

	//dec := json.NewDecoder(bytes.NewBuffer(jsonData))
	//dec.UseNumber() //关键步骤
	//err := dec.Decode(&v)

	if err != nil {
		return nil, err
	}
	data := v.(map[string]interface{})

	return data, nil
}

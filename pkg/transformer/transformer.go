package transformer

import "encoding/json"

func EncodeDecode[T any](i interface{}) (*T, error) {
	rawData, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var t T
	err = json.Unmarshal(rawData, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

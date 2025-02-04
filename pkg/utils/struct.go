package utils

import "encoding/json"

func StructToStruct[T, U any](t *T) (*U, error) {
	var u U
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

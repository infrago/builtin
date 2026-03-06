package builtin

import (
	"encoding/json"

	"github.com/infrago/infra"
)

func EncryptDIGIT(v int64) (string, error) {
	return infra.EncodeInt64(v)
}

func EncryptDIGITS(v []int64) (string, error) {
	return infra.EncodeInt64Slice(v)
}

func DecryptDIGIT(v string) (int64, error) {
	return infra.DecodeInt64(v)
}

func DecryptDIGITS(v string) ([]int64, error) {
	values, err := infra.DecodeInt64Slice(v)
	if err == nil {
		return values, nil
	}
	single, singleErr := infra.DecodeInt64(v)
	if singleErr != nil {
		return nil, err
	}
	return []int64{single}, nil
}

func EncryptTEXT(v string) (string, error) {
	return infra.EncodeTextBytes([]byte(v))
}

func EncryptTEXTS(v []string) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return infra.EncodeTextBytes(data)
}

func DecryptTEXT(v string) (string, error) {
	data, err := infra.DecodeTextBytes(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecryptTEXTS(v string) ([]string, error) {
	data, err := infra.DecodeTextBytes(v)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []string{}, nil
	}

	var values []string
	if err := json.Unmarshal(data, &values); err == nil {
		return values, nil
	}

	return []string{string(data)}, nil
}

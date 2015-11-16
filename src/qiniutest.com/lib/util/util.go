package util

import (
	"encoding/base64"
	"encoding/json"
	"github.com/jmcvetta/randutil"
	"os"
)

// ---------------------------------------------------------------------------

const (
	EncodeType_Base64    = iota
	EncodeType_Base64URL = iota
)

func Encode(encodeType int, Data string) (string, error) {
	switch encodeType {
	case EncodeType_Base64:
		return base64.StdEncoding.EncodeToString([]byte(Data)), nil
	case EncodeType_Base64URL:
		return base64.URLEncoding.EncodeToString([]byte(Data)), nil
	default:
		return "", nil
	}
}

func Decode(encodeType int, Data string) (string, error) {
	switch encodeType {
	case EncodeType_Base64:
		ret, err := base64.StdEncoding.DecodeString(Data)
		if err != nil {
			return "Decode ERROR", err
		}
		return string(ret), nil
	case EncodeType_Base64URL:
		ret, err := base64.URLEncoding.DecodeString(Data)
		if err != nil {
			return "Decode ERROR", err
		}
		return string(ret), nil
	default:
		return "", nil
	}
}

func ConvertToMap(data string) (map[string]interface{}, error) {
	b := []byte(data)
	var r interface{}
	err := json.Unmarshal(b, &r)
	return r.(map[string]interface{}), err
}

func MapCopy(dst, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}

func GetEnv() map[string]interface{} {
	test_env := os.Getenv("TEST_ENV")
	test_zone := os.Getenv("TEST_ZONE")
	envMap, _ := ConvertToMap(os.Getenv("QiniuTestEnv_" + test_env))
	zoneMap, _ := ConvertToMap(os.Getenv("QiniuTestEnv_" + test_env + "_" + test_zone))
	MapCopy(envMap, zoneMap)
	return envMap
}

func GetRand(n int) (str string) {
	str, _ = randutil.String(8, randutil.Alphanumeric)
	return
}

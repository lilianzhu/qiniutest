package util

import (
	"encoding/base64"
	"encoding/json"
	"github.com/jmcvetta/randutil"
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

func EncodeURL(url string) (res string) {
	res, _ = Encode(EncodeType_Base64URL, url)
	return res
}

func StringToJson(data string, r interface{}) error {
	b := []byte(data)
	return json.Unmarshal(b, &r)
}

func MapCopy(dst, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}

func GetRand(n int) (str string) {
	str, _ = randutil.String(n, randutil.Alphanumeric)
	return
}

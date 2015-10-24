package util

import (
	"encoding/base64"
)

// ---------------------------------------------------------------------------

const (
	EncodeType_Base64 = iota
	EncodeType_Base64URL = iota
)

func Qiniu_base64(IsStdEncoding bool,IsDecode bool,Data string)(string, error) {
	encoding := base64.URLEncoding
	if IsStdEncoding {
		encoding = base64.StdEncoding
	}
	if IsDecode {
		ret, err := encoding.DecodeString(Data)
		if err != nil {
			return "Decode ERROR", err
		}
		return string(ret), nil
	} else {
		return encoding.EncodeToString([]byte(Data)), nil
	}
}

func Encode(encodeType int, Data string)(string, error){
	switch encodeType {
		case EncodeType_Base64:
			return base64.StdEncoding.EncodeToString([]byte(Data)), nil
		case EncodeType_Base64URL:
			return base64.URLEncoding.EncodeToString([]byte(Data)), nil
		default:
			return "", nil
	}
}

func Decode(encodeType int, Data string)(string, error){
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

